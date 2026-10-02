package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	// ensure a name was passed as an arg
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no name provided")
	}

	// Create a new user in the database
	// Have access to the CreateUser query through state -> db struct
	// Pass context.Background() to the query to create an empty context argument

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	_, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		//fmt.Printf("Error creating user: %v\n", err)
		os.Exit(1)
	}

	// On Success update current user in config and print a success message for debugging
	s.config.SetUser(cmd.Args[0])
	fmt.Printf("User created successfully: %s\n ID: %s\n Created At: %s\n Updated At: %s\n", userParams.Name, userParams.ID, userParams.CreatedAt, userParams.UpdatedAt)
	return nil
}
