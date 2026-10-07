package main

import (
	"context"
	"fmt"

	"github.com/DorothyDanger/BlogAggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		userN, err := s.db.GetUser(context.Background(), s.config.Current_User_Name)
		if err != nil {
			return fmt.Errorf("failed to get current user: %v", err)
		}
		return handler(s, cmd, userN)
	}
}
