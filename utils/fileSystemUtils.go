package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Creates temporary directories using the given tree
// source: https://pkg.go.dev/path/filepath#Walk
func PrepareTestDirTree(tree string) (string, error) {
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

// Prints names of all entires in the current directory
// source: https://pkg.go.dev/path/filepath#Walk
func PrintAllFilenames() {
	// create sample directory entries
	tmpDir, err := PrepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}

	// delete sample directory entries at end of function
	defer os.RemoveAll(tmpDir)

	// enter new sample directory root
	os.Chdir(tmpDir)

	// list which entries to skip
	subDirToSkip := "skip"

	// traverse all directory items except the skipped one, and print their names
	fmt.Println("On Unix:")
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// if directory entry matches skip condition, return the SkipDir "error"
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}

		// print directory entry name
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})

	// handle filepath.Walk() error
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
}
