package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var supportCommand map[string]cliCommand

func init() {
	supportCommand = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
