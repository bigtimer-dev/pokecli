package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
		printHelperArea(&resp, cfg)
		return nil
	}
	resp, raw, err := cfg.client.ListLocations(cfg.next)
	if err != nil {
		return fmt.Errorf("error getting response from request : %w ", err)
	}
	cfg.cache.Add(key, raw)

	printHelperArea(&resp, cfg)
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
		printHelperArea(&resp, cfg)
		return nil
	}

	resp, raw, err := cfg.client.ListLocations(cfg.previous)
	if err != nil {
		return fmt.Errorf("error getting response from request : %w ", err)
	}

	cfg.cache.Add(key, raw)

	printHelperArea(&resp, cfg)
	return nil
}

func commandClear(cfg *config, mystring []string) error {
	fmt.Print("\033[H\033[2J")
	return nil
}

func commandExplore(cfg *config, mystring []string) error {
	if len(mystring) < 2 || len(mystring) > 2 {
		return fmt.Errorf("to use <explore> location-area")
	}
	key := mystring[1]
	if data, ok := cfg.cache.Get(key); ok {
		var resp pokeapi.PokemonInArea
		if err := json.Unmarshal(data, &resp); err != nil {
			return fmt.Errorf("error decoding entry: %w", err)
		}
		printHelperPokemon(&resp, cfg, key)
		return nil
	}
	resp, rawData, err := cfg.client.ExploreLocation(key)
	if err != nil {
		return err
	}
	cfg.cache.Add(key, rawData)

	printHelperPokemon(&resp, cfg, key)
	return nil
}

func commandCatch(cfg *config, mystring []string) error {
	if len(mystring) < 2 || len(mystring) > 2 {
		return fmt.Errorf("to use <catch> pokemon")
	}
	var resp pokeapi.Pokemon
	var rawData []byte
	key := mystring[1]

	if data, ok := cfg.cache.Get(key); ok {
		if err := json.Unmarshal(data, &resp); err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}
	} else {
		var err error

		resp, rawData, err = cfg.client.CaughtPokemon(key)
		if err != nil {
			return fmt.Errorf("error getting creature: %w", err)
		}
		cfg.cache.Add(key, rawData)
	}
	printCatch(resp, cfg, key)
	return nil
}

func printHelperArea(resp *pokeapi.Locations, cfg *config) {
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	cfg.next = resp.Next
	cfg.previous = resp.Previous
}

func printHelperPokemon(resp *pokeapi.PokemonInArea, cfg *config, key string) {
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", key)
	for _, Encounters := range resp.PokemonEncounters {
		fmt.Println(Encounters.Pokemon.Name)
	}
}

func tryingCatch(exp int) bool {
	max := exp
	if max < 20 {
		max = 20
	} else if max > 200 {
		max = 200
	}

	r := rand.Intn(max)
	return r < 50
}

func printCatch(resp pokeapi.Pokemon, cfg *config, key string) {
	if tryingCatch(resp.BaseExperience) {
		fmt.Printf("Throwing a Pokeball at %s...", key)
		fmt.Printf("%s was caught!", key)
		cfg.user.Add(key, resp)
	} else {
		fmt.Printf("Throwing a Pokeball at %s...", key)
		fmt.Printf("%s escaped!", key)
	}
}
