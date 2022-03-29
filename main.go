package main

import (
	"fmt"
	"os"
)

// searches all files in a directory tree for instances of a specified text
func main() {
	// test
	userArgs := os.Args[1:]

	// validate args
	if len(userArgs) < 1 {
		fmt.Println("Please provide an argument as in\n\tgo run . textToMatch")
		return
	}

	// test print arg
	fmt.Println("Finding all occurrences of \"" + userArgs[0] + "\" in current directory...")
}
