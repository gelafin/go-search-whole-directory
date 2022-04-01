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
	fmt.Println("Finding all occurrences of \"" + searchTerm + "\" in current directory...")

	// print all directory entries' paths
	utils.PrintAllFilepaths()

	utils.CountOccurrencesInCurrentDirectory(searchTerm)

	// print done message
	fmt.Print("\n" + DoneMessage)
}
