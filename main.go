package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bigtimer-dev/pokecli/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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

		fmt.Printf("Your command was: %s\n", slice[0])

		if err := scanner.Err(); err != nil {
			fmt.Fprint(os.Stderr, "Error scanning: ", err)
		}
	}
}
