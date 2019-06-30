package main

import (
	"fmt"
)

type Matcher struct {
	Name string `json:"site"`
	Count int
}

var matchers = make(map[string]Matcher)

// init methods in main are automagically called
// init methods are parameterless and have no return value
func init () {
	fmt.Println("initializing main")
	fmt.Println(matchers)
    matchers["rss"] = Matcher {Name: "rss", Count: 1}
	matchers["blogs"] = Matcher {"blogs", 1}
	matchers["default"] = Matcher {"default", 1}
}

func main () {
	fmt.Println(matchers)
}

