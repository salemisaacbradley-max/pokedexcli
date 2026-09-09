package main

import ("fmt"
		"os"
		"io"
		"net/http"
		"encoding/json")

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
		"map": {
		name: 		"map",
		description:"Prints the next 20 locations",
		callback:	commandMap,
	},
		"mapb": {
		name:		"mapb",
		description:"Prints the previous 20 locations",
		callback:	commandMapB,
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
	commands map[string]cliCommand
	previousLocationURL *string
	nextLocationURL *string
	pokecache Cache
}

type Location struct {
	Count int
	Next  *string
	Previous *string
	Results []struct {
		Name string
		URL string
	}
}

func commandMap (cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.nextLocationURL != nil {
		url = *cfg.nextLocationURL
	} 
	err = json.Unmarshal(body, &location)
	for _, i := range location.Results {
		fmt.Println(i.Name)
	}
	cfg.nextLocationURL = location.Next
	cfg.previousLocationURL = location.Previous
	return nil
}

func commandMapB (cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.previousLocationURL == nil {
		fmt.Println("you're on the first page")
	} else {
	}
	err = json.Unmarshal(body, &location)
	for _, i := range location.Results {
		fmt.Println(i.Name)
	}
	cfg.previousLocationURL = location.Previous
	cfg.nextLocationURL = location.Next}
	return nil
}

func processBytes (cache pokecache.Cache, url string) ([]byte, error) {
	
	cache.Get(url)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error with Map Response: %v\n", err)
	}
	location := Location{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving body: %v\n", err)
	}
	
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Unexpected Response Code: %v\n", res.StatusCode)
	}
	if err != nil {
		return nil, fmt.Errorf("Error with Body Response: %v\n", err)
	}
	return body, nil
}