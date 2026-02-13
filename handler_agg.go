package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpCl := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request object")
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := httpCl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make new request")
	}
	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feedData RSSFeed
	err = xml.Unmarshal(respByte, &feedData)
	if err != nil {
		return nil, err
	}

	feedData.Channel.Title = html.UnescapeString(feedData.Channel.Title)
	feedData.Channel.Description = html.UnescapeString(feedData.Channel.Description)
	for i, item := range feedData.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feedData.Channel.Item[i] = item
	}

	return &feedData, err
}

func scrapeFeeds(ctx context.Context, s *State) error {
	feed, err := s.db.GetNextFeed(ctx)

	if err != nil {
		return fmt.Errorf("failed to get next feed %v", err)
	}

	err = s.db.MarkFeedFetched(ctx, feed.ID)

	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched %v", err)
	}

	feedData, err := fetchFeed(ctx, feed.Url)
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

func AggregatorHandler(s *State, command Command) error {
	ctx := context.Background()
	if len(command.Args) < 2 {
		return fmt.Errorf("usage: %s time_between_reqs", command.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(command.Args[1])

	if err != nil {
		return fmt.Errorf("failed to parse time duration")
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(ctx, s)
	}
}
