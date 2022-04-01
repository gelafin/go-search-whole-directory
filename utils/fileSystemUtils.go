package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/tools/godoc/util"
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

func ShowOccurrencesInCurrentDirectory(searchTerm string) {
	previewCharCount := "10"

	// make regex including search term plus some surrounding chars.
	// Example: .{0,10}kindness.{0,10}
	searchTermRegEx := regexp.MustCompile(`.{0,` + previewCharCount + `}` + searchTerm + `.{0,` + previewCharCount + `}`)

	matchesInDir, err := findOccurrencesInCurrentDirectory(searchTermRegEx)

	// check for error
	if err != nil {
		return
	}

	// print all results
	for filepath, matchesInFile := range matchesInDir {
		fmt.Printf("in %v: ", filepath)
		for _, matchInBytes := range matchesInFile {
			fmt.Printf("\n\t%s\n", matchInBytes)
		}
	}
}

// Counts the occurrences of a given string in every file in the current directory,
// not including file and directory names,
// and returns results as a map, where keys are filepaths that were searched,
// and the value is the matched search term plus its immediate surroundings (10 chars in each direction)
func findOccurrencesInCurrentDirectory(searchTermRegEx *regexp.Regexp) (map[string][][]byte, error) {
	// prepare return variable
	matchesInDir := make(map[string][][]byte)

	// get a list of all files and folders to check
	allPaths, err := getAllDirEntries()

	// handle error in getting paths
	if err != nil {
		return matchesInDir, err
	}

	// find matches in all current directory files
	for _, path := range allPaths {
		// get matches in this file
		matchesInFile, err := findOccurrencesInFile(path, searchTermRegEx)

		// skip errored files and add valid counts to the total
		if err == nil && len(matchesInFile) > 0 {
			matchesInDir[path] = matchesInFile
		}
	}

	return matchesInDir, nil
}

// Finds the occurrences of a given string in a file,
// and returns the filename and a list of the matches
func findOccurrencesInFile(filepath string, searchTerm *regexp.Regexp) ([][]byte, error) {
	// prepare return variables
	matchesInFile := [][]byte{}

	// open the file
	fileData, err := os.ReadFile(filepath)

	// handle file open error by returning error to caller
	if err != nil {
		return matchesInFile, err
	}

	// For text files only,
	// search the file for all matches
	if util.IsText(fileData) {
		// fmt.Println("\tDEBUG: looks like text file!")
		matchesInFile = append(matchesInFile, searchTerm.FindAll(fileData, -1)...)
	}

	// return the total count
	return matchesInFile, nil
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
	var dirEntryPaths []string
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		// handle error
		if err != nil {
			fmt.Printf("failed to access %q: %v\n", path, err)
			return err
		}

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
func PrintAllFilepaths() {
	dirEntryPaths, err := getAllDirEntries()

	if err != nil {
		fmt.Println("error getting all directory names")
	}

	fmt.Print("printing all directory names\n\n")
	for _, path := range dirEntryPaths {
		fmt.Println(path)
	}
}
