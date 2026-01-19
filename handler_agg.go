package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

const FEED_URL = "https://www.wagslane.dev/index.xml"

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	req.Header.Set("User-Agent", "gator")

	if err != nil {
		return nil, fmt.Errorf("failed to create request object")
	}

	httpCl := http.Client{}
	resp, err := httpCl.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to make new request")
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	feedData := RSSFeed{}

	xml.Unmarshal(respByte, &feedData)

	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}

	feedData.Channel.Title = html.UnescapeString(feedData.Channel.Title)
	feedData.Channel.Description = html.UnescapeString(feedData.Channel.Description)

	return &feedData, err
}

func AggregatorHandler(s *State, command Command) error {
	ctx := context.Background()

	feedData, err := fetchFeed(ctx, FEED_URL)
	if err != nil {
		return fmt.Errorf("failed to fetch feeds")
	}

	for _, item := range feedData.Channel.Item {
		if item.Title == "Optimize For Simplicity First" {
			fmt.Print("Optimize for simplicity \n")
		} else {
			fmt.Printf("%s \n", item.Title)
		}
	}

	return nil
}
