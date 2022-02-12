package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type message struct {
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
