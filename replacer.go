package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"os"
  "encoding/json"
)

type Value struct{
  Source string
  value string
}

type Filter struct{
  Type string
  filter string
}

type Config struct{
  Brand string
  Token string
  DefaultDirection string
  Values []Value
}

func main() {
	// command line args without the prog
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: replacer {input options file} {replacement path}")
		os.Exit(0)
	}

	options := args[0]
	//source := args[1]

	// check if the options file exists
	optionsFile, err := ioutil.ReadFile(options)
	if err != nil {
		fmt.Println("Error reading input file: " + options)
		panic(err)
	}

  // load the options file
  var conf Config
  err = json.Unmarshal(optionsFile, &conf)
  if err != nil { panic(err) }

	// is the source a file or directory?
	// if its a file, lets start the replacements

	// ok, its a directory.  Run the replacement on all the files in it

}
