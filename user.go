package main

import "github.com/bigtimer-dev/pokecli/pokeapi"

type User struct {
	creatures map[string]pokeapi.Pokemon
}

func NewUser() *User {
	u := &User{
		creatures: make(map[string]pokeapi.Pokemon),
	}
	return u
}

func (c *User) Add(key string, val pokeapi.Pokemon) {
	c.creatures[key] = val
	return
}
