package main

import (
	// fmt is part of the standard library
	// standard library imports only need
	// to reference the package compared to
	// non standard libraries
	"fmt"

	// a blank import is used to initialize bank package
	// it will call any init functions in this package
	// and since we don't need anything else in bank, it
	// is discarded by _
	_ "github.com/mygoinaction/practice/bank"

	// the compiler will always look for
	// imports in GOROOT and GOPATH
)

func main() {
	fmt.Println("Main function")
}
