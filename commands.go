package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for key, value := range supportCommand {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	resp, err := cfg.client.ListLocations(cfg.next)
	if err != nil {
		return err
	}
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	cfg.next = resp.Next
	cfg.previous = resp.Previous
	return nil
}

func commandMapb(cfg *config) error {
	resp, err := cfg.client.ListLocations(cfg.previous)
	if err != nil {
		return err
	}
	if cfg.previous == nil {
		fmt.Println("You are on the first 20 locations")
	}
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	cfg.next = resp.Next
	cfg.previous = resp.Previous
	return nil
}
