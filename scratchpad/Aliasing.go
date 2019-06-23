package main

import (
	"fmt"
	s "strings"
)

// you can alias a function by assigning it to a variable
// because functions are 1st class constructs in Go
var p = fmt.Println

// you can alias a builtin type
// by redeclaring it
type MyOwnStuff float64

func Abs(mos MyOwnStuff) float64 {
	if (mos < 0) {
		return float64(0 - mos)
	}
	return float64(mos)
}

func main () {
	p("To Upper ", s.ToUpper("simple"))
	p("To Lower ", s.ToLower("SIMPLE"))

	var test float64
	var duration MyOwnStuff

	test = 54.0
	duration = 33.3
	p("A float64", test)
	p("A float64", duration)
}
