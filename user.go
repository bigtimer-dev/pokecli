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
}

func (c *User) Get(key string) (pokeapi.Pokemon, bool) {
	data, ok := c.creatures[key]
	return data, ok
}

func (c *User) All() []string {
	slice := []string{}
	for key := range c.creatures {
		slice = append(slice, key)
	}
	return slice
}
