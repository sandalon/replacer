package main

import (
	"./replacer"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// command line args without the prog
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: replacer {input options file} {replacement path}")
		return
	}

	options := args[0]
	destination := args[1]

	// check if the options file exists
	optionsFile, err := ioutil.ReadFile(options)
	if err != nil {
		fmt.Println("Error reading input file: " + options)
		return
	}

	// load the options file
	replacer.LoadOptions(optionsFile)

	// walk the file system from the destination
	err = filepath.Walk(destination, visit)
	if err != nil {
		fmt.Println("Error while replacing")
		panic(err)
	}
}

func visit(path string, fi os.FileInfo, err error) error {
	switch mode := fi.Mode(); {
	case mode.IsDir():
		if replacer.IsValidDirectory(fi.Name()) != true {
			return filepath.SkipDir
		}

	case mode.IsRegular():
		if replacer.IsValidFile(fi.Name()) {
			replacer.ProcessFile(path)
		}
	}

	return nil
}
