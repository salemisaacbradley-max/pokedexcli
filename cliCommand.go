package main

import ("fmt"
		"os"
		"io"
		"net/http")

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
	next map[string]cliCommand
	previous map[string]cliCommand
}

func commandMap (cfg *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/{id or name}/")
	if err != nil {
		return fmt.Errorf("Error with Map Response: %v\n", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Unexpected Response Code: %v\n", res.StatusCode)
	}
	if err != nil {
		return fmt.Errorf("Error with Body Response: %v\n", err)
	}
	return nil

}

func commandMapB (cfg *config) error {

}