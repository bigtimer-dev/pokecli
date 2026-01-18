package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bigtimer-dev/pokecli/pokeapi"
)

func commandExit(cfg *config, mystring []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, mystring []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range supportCommand {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(cfg *config, mystring []string) error {
	key := "https://pokeapi.co/api/v2/location-area"
	if cfg.next != nil {
		key = *cfg.next
	}
	if entry, ok := cfg.cache.Get(key); ok {
		var resp pokeapi.Locations
		if err := json.Unmarshal(entry, &resp); err != nil {
			return fmt.Errorf("error decoding entry: %w", err)
		}
		printHelper(&resp, cfg)
		return nil
	}
	resp, raw, err := cfg.client.ListLocations(cfg.next)
	if err != nil {
		return fmt.Errorf("error getting response from request : %w ", err)
	}
	cfg.cache.Add(key, raw)

	printHelper(&resp, cfg)
	return nil
}

func commandMapb(cfg *config, mystring []string) error {
	if cfg.previous == nil {
		fmt.Println("You are on the first 20 locations")
		return nil
	}

	key := *cfg.previous

	if data, ok := cfg.cache.Get(key); ok {
		var resp pokeapi.Locations
		if err := json.Unmarshal(data, &resp); err != nil {
			return fmt.Errorf("error decoding entry: %w", err)
		}
		printHelper(&resp, cfg)
		return nil
	}

	resp, raw, err := cfg.client.ListLocations(cfg.previous)
	if err != nil {
		return fmt.Errorf("error getting response from request : %w ", err)
	}

	cfg.cache.Add(key, raw)

	printHelper(&resp, cfg)
	return nil
}

func commandClear(cfg *config, mystring []string) error {
	fmt.Print("\033[H\033[2J")
	return nil
}

func commandExplore(cfg *config, mystring []string) error {
}

func printHelper(resp *pokeapi.Locations, cfg *config) {
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	cfg.next = resp.Next
	cfg.previous = resp.Previous
}
