package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bigtimer-dev/pokecli/pokeapi"
	"github.com/bigtimer-dev/pokecli/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		client: pokeapi.NewClient(),
	}
	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		slice := repl.CleanInput(line)

		if len(slice) == 0 {
			continue
		}

		if value, ok := supportCommand[slice[0]]; ok {
			if err := value.callback(cfg); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknow command")
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprint(os.Stderr, "Error scanning: ", err)
		}
	}
}
