package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Language    string    `xml:"language"`
	Item        []RssItem `xml:"item"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// fetchRSS fetches RSS feed from the given URL and returns a Rss struct and an error.
func fetchRSS(url string) (Rss, error) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	
	if err != nil {
		return Rss{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return Rss{}, err
	}

	rssFeed := Rss{}

	err = xml.Unmarshal(data, &rssFeed)

	if err != nil {
		return Rss{}, err
	}

	return rssFeed, nil

}
