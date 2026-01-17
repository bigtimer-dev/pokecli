package repl

import "strings"

func CleanInput(mystring string) []string {
	newString := strings.ToLower(mystring)
	slice := strings.Fields(newString)
	return slice
}
