package repl

import "strings"

func CleanInput(mystring string) []string {
	slice := strings.Fields(mystring)
	return slice
}
