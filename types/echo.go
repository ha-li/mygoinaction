package main

import (
	"io"
	"os"
)

func main() {

	// a simple program that implements echo continuous
	// no need for the for loop

	// for {
	io.Copy (os.Stdout, os.Stdin)
	//}

	// os.Stdin implements io.Reader
	// os.Stdout implements io.Writer

}