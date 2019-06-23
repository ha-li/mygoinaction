package main

import "fmt"

func main () {
	// the ellipsis can be used to dynamically allocate
	// an array
	var c = [...]string {
		"abbc",
		"defe",
		"high",
		"jkl",
	}

	fmt.Println(c);

	// you cannot allocate an
	// array of certain size and then
	// over index it, causes an
	// index out of bound error
	var d = [3]string {
		"abc",
		"def",
		"high",
		//"lmnl"
	}

	// use %q to print each array element quoted
	fmt.Printf("%q\n", d)

	is := [...]int {
		1, 2, 3, 4, 5, 6,
	}

	fmt.Printf("%q\n", is)
}
