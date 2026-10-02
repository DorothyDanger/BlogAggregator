package main

import (
	"context"
	"fmt"
)

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
