package main

import ("fmt"
		"os")

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

var commands map[string]cliCommand

func initCommands() {
    commands = map[string]cliCommand{
        "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
		"help": {
		name:		"help",
		description:"Displays a help message",
		callback:	commandHelp,
	},
    }
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	for _, comm := range commands {
		fmt.Printf("%s: %s\n", comm.name, comm.description)
	}
	return nil
}

type config struct {
	chart map[string]cliCommand
}