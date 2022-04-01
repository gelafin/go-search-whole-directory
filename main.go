package main

import (
	"fmt"
	"os"

	"github.com/gelafin/go-search-whole-directory/utils"
)

// searches all files in a directory tree for instances of a specified text
func main() {
	// get command-line args
	userArgs := os.Args[1:]

	// validate necessary args were passed
	if len(userArgs) < 1 {
		fmt.Println(ErrorMessages[MissingCommandLineArg])
		return
	}

	searchTerm := userArgs[0]

	// print start message
	fmt.Print("Finding all occurrences of \"" + searchTerm + "\" in current directory...\n\n")

	utils.ShowOccurrencesInCurrentDirectory(searchTerm)

	// print done message
	fmt.Print("\n" + DoneMessage)
}
