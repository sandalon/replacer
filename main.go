package main

import (
	"fmt"
	"./replacer"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// command line args without the prog
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: replacer {input options file} {replacement path}")
		os.Exit(0)
	}

	options := args[0]
	source := args[1]

	// check if the options file exists
	optionsFile, err := ioutil.ReadFile(options)
	if err != nil {
		fmt.Println("Error reading input file: " + options)
		return
	}

	// load the options file
	replacer.LoadOptions(optionsFile)

	err = filepath.Walk(source, visit)
	if err != nil {
		panic(err)
	}
}


func visit(path string, fi os.FileInfo, err error) error {

	switch mode := fi.Mode(); {
	case mode.IsRegular(): {
		fmt.Println("Opening file: " + path)
	}
	}

	return nil
}
