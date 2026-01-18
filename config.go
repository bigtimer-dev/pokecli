package main

import "github.com/bigtimer-dev/pokecli/pokeapi"

type config struct {
	next     *string
	previous *string
	client   *pokeapi.Client
}
