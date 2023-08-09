package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type message struct {
	account             account
	content, visibility string
	feed                feed
}

func (msg *message) post() error {
	if debug {
		log.Printf("msg:\n\n%v", msg)
		log.Printf("msg.feed:\n\n%v", msg.feed)
	}

	apiURL := msg.account.InstanceURL + "/api/v1/statuses"

	data := url.Values{}
	data.Set("status", msg.content)
	data.Set("visibility", msg.feed.Visibility)
	data.Set("sensitive", strconv.FormatBool(msg.feed.Sensitive))
	data.Set("spoiler_text", msg.feed.ContentWarning)

	if debug {
		log.Printf("Data:\n\n%v", data)
	}

	switch msg.feed.Format {
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
		log.Printf("Message:\n\n%s", msg.content)
	}

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return err
	}

	if debug {
		log.Printf("Request:\n\n%v", body)
	}

	// Set Headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+msg.account.AccessToken)
	req.Header.Set("User-Agent", yamlfile.HttpConfig.UserAgent)

	c := &http.Client{Timeout: time.Second * 10}

	if !debugOrDryRun {
		var resp *http.Response
		resp, err = c.Do(req)
		if err != nil {
			return err
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		log.Println("response Status: ", resp.Status)
	}

	return nil
}
