package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"

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

	var tpls = make(map[string]*template.Template)
	var formats = make(map[string]string)
	for account_index, a := range conf.Accounts {
		for feed_index, f := range a.Feeds {
			tmpl, err := template.New(f.URL).Parse(f.Template)
			if err != nil {
				log.Fatalf("Failed to parse template [%s]. Error: %s", f.Template, err.Error())
			}
			// Default format to "plain", if blank
			if f.Format == "" {
				conf.Accounts[account_index].Feeds[feed_index].Format = "plain"
				f.Format = "plain"
			}
			tpls[f.URL] = tmpl
			formats[f.URL] = f.Format
		}
	}

	for _, account := range conf.Accounts {
		// Get feeds
		log.Printf("Fetching feeds for account [%s]...", account.Name)
		var feeds []*rss.Feed
		for _, source := range account.Feeds {
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
				if item.Date.Before(conf.LastUpdated) {
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
					log.Fatalf("Error executing template [%s]. Error: %s", tpls[base.String()], err.Error())
				}
				if err = postMessage(account, buf.String(), formats[base.String()]); err != nil {
					log.Fatalf("Failed to post message \"%s\". Error: %s", buf.String(), err.Error())
				}
			}
		}
	}

	if !debugOrDryRun {
		// update timestamp in config
		conf.updateLastUpdated()
		// save config
		conf.Save()
	}
}

func postMessage(account Account, message string, format string) error {
	apiURL := account.InstanceURL + "/api/v1/statuses"

	data := url.Values{}
	data.Set("status", message)
	data.Set("visibility", "unlisted")

	switch format {
	case "markdown":
		data.Set("content_type", "text/markdown")
	case "html":
		data.Set("content_type", "text/html")
	case "plain":
		data.Set("content_type", "text/plain")
	case "bbcode":
		data.Set("content_type", "text/bbcode")
	}

	var req *http.Request
	var body io.Reader
	body = strings.NewReader(data.Encode())

	if debug {
		log.Printf("Message:\n\n%s", message)
	}

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return err
	}

	// Set Headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+account.AccessToken)

	c := &http.Client{Timeout: time.Second * 10}

	if !debugOrDryRun {
		var resp *http.Response
		resp, err = c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		log.Println("response Status: ", resp.Status)
	}

	return nil
}
