package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
	"github.com/google/uuid"
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

	for _, item := range rss.Channel.Item {
		pubDate := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			pubDate = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}
		_, err := s.db.CreatePost(ctx, postParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("Error: %v", err)
			continue
		}
		fmt.Printf("Added post to the database for user")
	}
	return nil
}
