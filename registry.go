package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"map": {
			name:        "map",
			description: "Displays the 20 next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the 20 preview locations",
			callback:    commandMapb,
		},
		"clear": {
			name:        "clear",
			description: "Clears the console screen",
			callback:    commandClear,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemons found in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "stats of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show all caught pokemons",
			callback:    commandPokedex,
		},
	}
}
