package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) >= 1 {
		if lim, err := strconv.Atoi(cmd.Args[0]); err != nil {
			return fmt.Errorf("invalid limit: %v\n Usage: limit <num>", err)
		} else {
			limit = lim
		}
	}

	userPostParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), userPostParams)
	if err != nil {
		return fmt.Errorf("Error fetching posts: %v", err)
	}
	fmt.Printf("Found %d posts for user %s\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Description: %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("Published at: %s\n\n", post.PublishedAt.Time)
	}
	return nil
}
