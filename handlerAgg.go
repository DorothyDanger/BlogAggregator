package main

import (
	"context"
	"fmt"
)

// later this will be long-running aggregator service.
// For now it will be used to fetch a single feed to ensure parsing works
// It should fetch the feed found at https://www.wagslane.dev/index.xml and print the entire struct to the console.
func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %v\n", err)

	}

	fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
	fmt.Printf("Feed Link: %s\n", feed.Channel.Link)
	fmt.Printf("Feed Description: %s\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Printf("\nItem Title: %s\n", item.Title)
		fmt.Printf("Item Link: %s\n", item.Link)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Printf("Item PubDate: %s\n", item.PubDate)
	}
	return nil
}
