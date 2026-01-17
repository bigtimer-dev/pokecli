package repl

import "strings"

func cleanInput(mystring string) []string {
	slice := strings.Fields(mystring)
	return slice
}
