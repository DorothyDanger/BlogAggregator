package main

import (
	"context"
	"fmt"
	"os"
)

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
