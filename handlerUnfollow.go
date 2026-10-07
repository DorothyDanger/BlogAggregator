package main

import (
	"context"
	"fmt"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	// Takes a single URL as an argument and unfollows it for the current user.
	if len(cmd.Args) < 1 {
		return fmt.Errorf("No URL provided. Usage: unfollow <url>\n")
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Failed to get feed from URL: %s\n Error: %v", url, err)
	}
	deleteParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFeedFollow(context.Background(), deleteParams)
	fmt.Printf("Successfully unfollowed feed: %s\n", feed.Name)
	fmt.Printf("Current user: %s\n", user.Name)
	return nil
}
