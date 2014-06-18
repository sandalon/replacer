package replacer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Value struct {
	Source string `json:"Source"`
	Value  string `json:"value"`
}

type Filter struct {
	Type          string `json:"Type"`
	FilterPattern string `json:"filterPattern"`
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
  isValid := true
	for _, filter := range conf.Filters {
		if filter.Type == "Directory" {
			splitDir := strings.Split(path, string(os.PathSeparator))
			for _, dir := range splitDir {
				if dir == filter.FilterPattern {
					isValid = false
				}
			}
		}
	}

	return isValid
}

func IsValidFile(path string) bool {
	return IsValidFileWithConf(path, conf)
}

func IsValidFileWithConf(path string, conf Config) bool {
	for _, filter := range conf.Filters {
		if filter.Type == "File" {
			extension := filepath.Ext(path)
			if extension == ("." + filter.FilterPattern) {
				return true
			}
		}
	}

	return false
}

func ProcessFile(path string) {
	fmt.Println("Processing file: " + path)
  replacementMade := false

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file: " + path)
		return
	}

  text := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    text += scanner.Text() + "\n"
	}
	file.Close()

  if strings.HasSuffix(text, "\n"){
      text = text[:len(text)-1]
  }

  switch conf.DefaultDirection {
  case "forward":
    for _, rule := range conf.Values {
      search := conf.Token + rule.Source + conf.Token
      replacement := rule.Value

      if(strings.Contains(text, search)){
        text = strings.Replace(text, search, replacement, -1)
        replacementMade = true
      }
    }
  case "reverse":
    for _, rule := range conf.Values {
      search := rule.Value
      replacement := conf.Token + rule.Source + conf.Token

      if(strings.Contains(text, search)){
        text = strings.Replace(text, search, replacement, -1)
        replacementMade = true
      }
    }
  }

  if replacementMade {
    file, err = os.Create(path)
    if err != nil {
      fmt.Println("Error creating file: " + path)
      return
    }

    fmt.Fprintln(file, text)

    file.Close()
  }
}
