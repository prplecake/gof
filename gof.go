package main

import (
	"log"
	"net/url"

	//"github.com/McKael/madon"
	"github.com/SlyMarbo/rss"
)

func main() {
	log.Println("gof starting up...")
	config := readConfig()

	// Get feeds
	log.Println("Fetching feeds...")
	var feeds []*rss.Feed
	for _, source := range config.Feeds {
		feed, err := rss.Fetch(source.URL)
		if err != nil {
			log.Printf("Error fetching %s: %s", source.URL, err.Error())
			continue
		}
		feeds = append(feeds, feed)
		log.Printf("Fetched %s", feed.Title)
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
			itemLink, err := url.Parse(item.Link)
			if err != nil {
				log.Fatal("failed parsing article URL of the feed item")
			}

			// Make sure data looks OK
			// TODO: remove before release
			log.Printf("Item Data:\n\tTimestamp: %s\n\tSite URL: %s\n\tFeed Title: %s\n\tItem Title: %s\n\tItem URL: %s\n",
				item.Date, base.ResolveReference(feedLink).String(),
				feed.Title, item.Title, base.ResolveReference(itemLink).String())
		}
	}

	// update timestamp in config
	config.updateLastUpdated()
	// save config
	config.Save()
}
