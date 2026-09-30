package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DorothyDanger/BlogAggregator/internal/config"
)

const configFileName = ".gatorconfig.json"

func main() {

	// read the config file and store it in cfg
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Failed to read config: %v\n", err)
		return
	}

	// Store the config in a new instance of the state struct

	s := &state{
		config: &cfg,
	}

	// Create a new instance of the commands struct
	// with an initialized map of handler functions
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	// Register a handler function for the login command
	cmds.register("login", handlerLogin)

	// Use os.Args to get the command-line arguments passed in by the user
	// If <2 print an error to the terminal and exit. 1st is command name, 2nd is input
	if len(os.Args) < 2 {
		log.Fatal("Error.")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(s, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
