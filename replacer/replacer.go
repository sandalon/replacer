package replacer

import (
  "fmt"
  "encoding/json"
)

type Value struct {
	Source string
	value  string
}

type Filter struct {
	Type   string
	filter string
}

type Config struct {
	Brand            string
	Token            string
	DefaultDirection string
	Values           []Value
}

var conf Config

func LoadOptions(file []byte) {
  err := json.Unmarshal(file, &conf)
  if err != nil {
    fmt.Println("Error parsing the options file")
    return
  }

  fmt.Println("Processing Brand: " + conf.Brand)
  fmt.Println("Process Direction: " + conf.DefaultDirection)
}

func ProcessFile(path string){

}
