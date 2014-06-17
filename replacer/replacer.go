package replacer

import (
	"encoding/json"
	"fmt"
  "strings"
  "os"
)

type Value struct {
	Source string
	value  string
}

type Filter struct {
	Type   string `json:"Type"`
	FilterPattern  string `json:"filterPattern"`
}

type Config struct {
	Brand            string
	Token            string
	DefaultDirection string
	Filters          []Filter
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

func IsValidDirectory(path string) bool {
  return IsValidDirectoryWithConf(path, conf)
}

func IsValidDirectoryWithConf(path string, conf Config) bool {
	for _, filter := range conf.Filters {
    if filter.Type == "Directory" {
      splitDir := strings.Split(path, string(os.PathSeparator))
      fmt.Println(path)
      fmt.Println(splitDir[0])
      for _, dir := range splitDir {
        if dir == filter.FilterPattern {
          return true
        }
      }
    }
	}

	return false
}

func IsValidFile(path string) bool {
  return IsValidFileWithConf(path, conf)
}

func IsValidFileWithConf(path string, conf Config) bool {
	return true
}

func ProcessFile(path string) {
	fmt.Println("Processing file: " + path)
}
