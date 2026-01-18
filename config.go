package main

import (
	"github.com/bigtimer-dev/pokecli/pokeapi"
	"github.com/bigtimer-dev/pokecli/pokecache"
)

type config struct {
	next     *string
	previous *string
	cache    *pokecache.Cache
	client   *pokeapi.Client
	user     *User
}
