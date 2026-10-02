package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	// takes 2 args: name of the feed, url of the feed
	// get current user from the database and connect the feed to that user.
	// print out the fields of the new feed record if successful
	if len(cmd.Args) < 2 {
		return fmt.Errorf("not enough arguments provided. Usage: addfeed <name> <url>\n")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// Get the current user ID from the database
	user, err := s.db.GetUser(context.Background(), s.config.Current_User_Name)
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	// Create a new feed in the database
	feedParams := database.AddFeedParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       url,
		UserID:    user.ID,
	}

	_, err = s.db.AddFeed(context.Background(), feedParams)
	if err != nil {
		//return fmt.Errorf("failed to add feed: %v", err)
		os.Exit(1)
	}

	// On Success print out the fields of the new feed record
	fmt.Printf("Feed Added Successfully\n")
	fmt.Printf("ID: %s\n", feedParams.ID)
	fmt.Printf("Name: %s\n", feedParams.Name)
	fmt.Printf("Created At: %s\n", feedParams.CreatedAt)
	fmt.Printf("Updated At: %s\n", feedParams.UpdatedAt)
	fmt.Printf("URL: %s\n", feedParams.Url)
	fmt.Printf("User ID: %s\n", feedParams.UserID)

	return nil
}
