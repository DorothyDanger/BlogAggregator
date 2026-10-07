package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	// Takes a single URL arg and creates a new feed follow record for the current user
	// Print the name of the feed and the current user once the record is created
	if len(cmd.Args) < 1 {
		return fmt.Errorf("No URL provided. Usage: follow <url>\n")
	}
	feedURL := cmd.Args[0]
	// Fetch feed data from database with the URL
	feed, err := s.db.GetFeedFromURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("Failed to get feed name from URL: %s\n Error: %v", feedURL, err)
	}

	// Create a new record in the database
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("failed to follow feed: %v", err)
	}
	fmt.Printf("Successfully followed feed: %s\n", feed.Name)
	fmt.Printf("Current user: %s\n", user.Name)
	return nil
}
