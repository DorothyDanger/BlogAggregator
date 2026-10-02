package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

type command struct {
	Name string
	Args []string
}

func handlerLogin(s *state, cmd command) error {
	// if command arg's slice is empty return an error
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username provided")
	}
	// Check if user exists in the database
	// exit with code 1 if user does not exist
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		os.Exit(1)
	}
	// use the state's access to the config struct to set user to the given username
	s.config.SetUser(cmd.Args[0])
	fmt.Printf("User set to: %s\n", s.config.Current_User_Name)
	return nil
}

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

func handlerReset(s *state, cmd command) error {
	// Resets the database by deleting all users
	// Calls ResetUsers query state -> db struct
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users: %v", err)
	}

	fmt.Println("Users reset successfully.")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	// Get all users from the database
	// Prints the list of users to the console

	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}
	for _, user := range users {
		if user.Name == s.config.Current_User_Name {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}
