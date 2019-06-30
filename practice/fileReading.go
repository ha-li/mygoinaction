package main

import (
	// for processing json stuff
	"encoding/json"
	"fmt"

	// for reading files
	"os"

	// don't really need this
	"path/filepath"
)

// in order to find this file, you need to run go run
// while inside the directory practice
const data = "data/tmp.json"

// a struct that is used to decode the json file data
// the 3rd construct on each property is a tag that
// maps the property to the json element in the file
type Element struct {
	Name   string `json:"key"`
	Symbol string `json:"value"`
}

func main() {
	filePath, err := filepath.Abs(data)
	// fmt.Println("abs path is ", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Errors while reading file %s\n errors %v", data, err)
	}

	// a built in that will schedule a function call once this context
	// finishes
	defer file.Close()

	// a slice of Elements
	var elements []Element
	err = json.NewDecoder(file).Decode(&elements)
	for _, element := range elements {
		fmt.Printf("%s %s\n", element.Name, element.Symbol)
	}
}