package main

import (
	"context"
	"fmt"
)

// This functions exists to be called by Agg.
// It is responsible for:
// 1.	Getting the next feed
// 2.	Marking it as fetched (updating last_fetched_at)
// 3.	Fetching the feed using the URL
// 4.	iterating over the items in the feed and printing the titles to the console
func scrapeFeed(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("Error finding next feed to fetch: %v\n", err)
	}

	err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		fmt.Errorf("Error marking feed as fetched: %v", err)
	}

	rss, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %v\n", err)
	}

	fmt.Printf("Feed Title: %s\n", rss.Channel.Title)
	fmt.Printf("Feed Link: %s\n", rss.Channel.Link)
	fmt.Printf("Feed Description: %s\n", rss.Channel.Description)
	for _, item := range rss.Channel.Item {
		fmt.Printf("\nItem Title: %s\n", item.Title)
		fmt.Printf("Item Link: %s\n", item.Link)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Printf("Item PubDate: %s\n", item.PubDate)
	}
	return nil
}
