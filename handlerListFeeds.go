package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func handlerListFeeds(s *state, cmd command) error {
	// No CLI args
	// prints all the feeds in the database to the console

	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}
	for _, feed := range feeds {
		fmt.Printf("Feed name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		err, userName := getUserNameByID(s, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Created by: %s\n", userName)
	}
	return nil
}

func getUserNameByID(s *state, userID uuid.UUID) (error, string) {
	// get user name from the database by user ID
	name, err := s.db.GetNameByID(context.Background(), userID)
	if err != nil {
		return fmt.Errorf("failed to get user name by ID: %v", err), ""
	}
	return nil, name
}
