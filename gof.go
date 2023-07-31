package main

import (
	"bytes"
	"flag"
	"log"
	"net/url"
	"text/template"

	"github.com/SlyMarbo/rss"
)

var (
	conf          *config
	debug         bool
	debugOrDryRun bool
)

type article struct {
	Title, URL, Summary string
}

func main() {
	var (
		configFile string
		dryRun     bool
	)
	flag.StringVar(&configFile, "c", "gof.yaml", "the configuration file to use")
	flag.BoolVar(&dryRun, "dry-run", false, "whether to perform a dry run or not")
	flag.BoolVar(&debug, "debug", false, "enable debugging")
	flag.Parse()

	debugOrDryRun = debug || dryRun

	log.Println(configFile)

	log.Println("gof starting up...")
	conf = readConfig(configFile)
	log.Printf("Version: %s\n", conf.Meta.Version)
	log.Printf("Build time: %s\n", conf.Meta.Buildtime)

	var tpls = make(map[string]*template.Template)
	var formats = make(map[string]string)
	for accountIndex, a := range conf.Accounts {
		for feedIndex, f := range a.Feeds {
			tmpl, err := template.New(f.URL).Parse(f.Template)
			if err != nil {
				log.Fatalf("Failed to parse template [%s]. Error: %s", f.Template, err.Error())
			}
			// Default format to "plain", if blank
			if f.Format == "" {
				conf.Accounts[accountIndex].Feeds[feedIndex].Format = "plain"
				f.Format = "plain"
			}
			tpls[f.URL] = tmpl
			formats[f.URL] = f.Format
		}
	}

	for _, account := range conf.Accounts {
		var toot message
		// Get feeds
		log.Printf("Fetching feeds for account [%s]...", account.Name)
		var feeds []*rss.Feed
		for _, source := range account.Feeds {
			toot.feed = source
			if debug {
				log.Printf("source:\n\n%v", source)
			}
			feed, err := rss.Fetch(source.URL)
			if err != nil {
				log.Printf("Error fetching %s: %s", source.URL, err.Error())
				continue
			}
			feeds = append(feeds, feed)
			log.Printf("Fetched [%s]", feed.Title)
		}
		if len(feeds) == 0 {
			log.Fatal("Expected at least one feed to successfully fetch.")
		}

		// Loop through feeds
		for _, feed := range feeds {
			// Get feed items
			if len(feed.Items) == 0 {
				log.Printf("Warning: feed %s has no items.", feed.Title)
				continue
			}
			items := feed.Items
			if len(items) > 1 {
				items = items[:1]
			}
			base, err := url.Parse(feed.UpdateURL)
			if err != nil {
				log.Fatal("failed parsing update URL of the feed")
			}
			feedLink, _ := url.Parse(feed.Link)
			if err != nil {
				log.Fatal("failed parsing canonical feed URL of the feed")
			}

			// Loop through items
			for _, item := range items {
				if item.Date.Before(conf.LastUpdated) && !debug {
					log.Println("No new items. Skipping.")
					continue
				}
				itemLink, err := url.Parse(item.Link)
				if err != nil {
					log.Fatal("failed parsing article URL of the feed item")
				}

				// Make sure data looks OK
				log.Printf("Item Data:\n\tTimestamp: %s\n\tSite URL: %s\n\tFeed Title: %s\n\tItem Title: %s\n\tItem URL: %s\n",
					item.Date, base.ResolveReference(feedLink).String(),
					feed.Title, item.Title, base.ResolveReference(itemLink).String())
				i := article{
					Title:   item.Title,
					Summary: item.Summary,
					URL:     base.ResolveReference(itemLink).String(),
				}
				buf := new(bytes.Buffer)
				err = tpls[base.String()].Execute(buf, i)
				if err != nil {
					log.Fatalf("Error executing template [%v]. Error: %s", tpls[base.String()], err.Error())
				}
				toot.account = account
				toot.content = buf.String()
				if err = toot.post(); err != nil {
					log.Fatalf("Failed to post message \"%s\". Error: %s", buf.String(), err.Error())
				}
			}
		}
	}

	if !debugOrDryRun {
		// update timestamp in config
		conf.updateLastUpdated()
		// save config
		err := conf.Save()
		if err != nil {
			log.Fatalf("Failed to save config to file. Error: %s", err.Error())
		}
	}
}
