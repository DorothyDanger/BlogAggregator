package main

import (
	"context"
	"fmt"
)

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
