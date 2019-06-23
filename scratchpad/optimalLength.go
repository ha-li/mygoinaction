package main

import (
	"fmt"
)


// the purpose of this exercise is to organize these
// names into arrays of names that are the same length
// the arrays should be stored in an array
// so the final datastructure is an array of arrays of strings
// or slices.
var names = []string {
	"Kaiser",
	"Sosa",
	"Moniche",
	"Tyrell",
	"Mike",
}

func main () {
	// make a map, maps are always great for instant lookup
	distribution := make(map[int][]string)

	// we'll need the max length (but we don't know this ahead of time
	// but it's easy to refactor after
	maxLength := 0

	// iterate through each of the names
	// O(n), where n = # of names
	for _, name := range names {
		// get the length of the name
		length := len(name)

		// save the max length if applicabble
		if length > maxLength {
			maxLength = length
		}

		// look for an array of holding the names of the same length
		// in the map, if not there, create a new slice
		// add to that slice
		if distribution[length] == nil {
			distribution[length] = make([]string, 1)
			distribution[length] = append(distribution[length], name)
		} else {
			distribution[length] = append(distribution[length], name)
		}
		//fmt.Println(len(distribution))
		fmt.Printf("%v\n", distribution)
	}

	// make an array of arrays from 0 to max length to hold
	// all the arrays
	// l := len(distribution)
	fin := make([][]string, maxLength)
	fmt.Println(fin)

	// get all the key pairs in the map
	// O(#len(n))
	for k,v := range distribution {
		// store them the arrays in their correct spot (len - 1) of the final array
		fin[k-1] = v
		//fmt.Printf ("%d, %v\n", k, v)
		//fin[k] = v
	}
	fmt.Printf("%v\n", fin)
}