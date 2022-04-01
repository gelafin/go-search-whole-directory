package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Creates temporary directories using the given tree
// source: https://pkg.go.dev/path/filepath#Walk
func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}

// Counts the occurrences of a given string in a file
func countOccurrencesInFile(filepath string, matchString string) {
	// open the file

	// search the file until finding the string

	// return the total count
}

// Adapted from https://pkg.go.dev/path/filepath#Walk
func getAllDirEntries() ([]string, error) {
	// create sample directory entries
	tmpDir, err := prepareTestDirTree("temp/sample")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return nil, err
	}

	// delete sample directory entries at end of function
	defer os.RemoveAll(tmpDir)

	// traverse all directory items, and store their names
	// DEBUG print every visited file path
	// fmt.Println("Assuming Unix paths:")
	var dirEntryPaths []string
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		// handle error
		if err != nil {
			fmt.Printf("failed to access %q: %v\n", path, err)
			return err
		}

		// DEBUG print visited file path
		// fmt.Printf("visited file or dir: %q\n", path)
		// append path to array kept in outer scope
		dirEntryPaths = append(dirEntryPaths, path)

		return nil
	})

	// handle filepath.Walk() error
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return nil, err
	}

	return dirEntryPaths, nil
}

// Prints names of all entires in the current directory
func PrintAllFilenames() {
	dirEntryPaths, err := getAllDirEntries()

	if err != nil {
		fmt.Println("error getting all directory names")
	}

	fmt.Print("printing all directory names\n\n")
	for _, path := range dirEntryPaths {
		fmt.Println(path)
	}
}
