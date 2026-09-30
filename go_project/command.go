package main

import (
	"fmt"
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
	// use the state's access to the config struct to set user to the given username
	s.config.SetUser(cmd.Args[0])
	fmt.Printf("User set to: %s\n", s.config.Current_User_Name)
	return nil
}
