package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const targetDir = "node_modules"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		handleError(errors.New("no directory path provided"))
	}
	dirPath := args[0]
	// check that dirpath exists
	fileInfo, err := os.Stat(dirPath)
	handleError(err)
	// check that dirpath is a directory
	if !fileInfo.IsDir() {
		handleError(errors.New(dirPath + " is not a directory"))
	}
	count := 0
	recRemoveNodeModules(dirPath, &count)
	fmt.Printf("Found and removed %d node_modules directories\n", count)
}

func handleError(err error) {
	if err == nil {
		return
	}
	_, writeErr := os.Stderr.WriteString(err.Error() + "\n")
	if writeErr != nil {
		panic(writeErr)
	}
	os.Exit(1)
}

func recRemoveNodeModules(dirPath string, count *int) {
	files, err := os.ReadDir(dirPath)
	handleError(err)
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		fullPath := path.Join(dirPath, file.Name())
		if file.Name() == targetDir {
			*count++
			println("removing " + fullPath)
			err := os.RemoveAll(fullPath)
			handleError(err)
		} else {
			recRemoveNodeModules(fullPath, count)
		}
	}
}
