package main

import (
	"fmt"
)

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// Runs a given command with the provided state if it exists
func (c *commands) run(s *state, cmd command) error {
	fn, exists := c.registeredCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("command not found: %s", cmd.Name)
	}

	return fn(s, cmd)
}

// registers a new handler function for a comamand name
func (c *commands) register(name string, f func(*state, command) error) {
	_, exists := c.registeredCommands[name]
	if exists {
		fmt.Printf("Command %s already has a handler registered:\n", name)
		return
	}
	c.registeredCommands[name] = f
}
