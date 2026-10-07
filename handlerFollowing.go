package main

import (
	"context"
	"fmt"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
)

func handlerPrintFollowedFeeds(s *state, cmd command, user database.User) error {
	// Function to print all the names of the feeds that the current user is following
	// No args are needed
	userID := user.ID

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), userID)
	if err != nil {
		return fmt.Errorf("failed to get followed feeds for user: %v", err)
	}

	fmt.Printf("Feeds followed be user %s:\n", s.config.Current_User_Name)
	for _, feed := range feeds {
		fmt.Printf("Feed name: %s\n", feed.FeedName)
	}
	return nil
}
