package main

// this illustrates closures

import (
	"fmt"
)

func ignored (msg string) {

	// this does not compile
	/* func() {
		fmt.Println(msg)
	} */

}

func runnable (msg string) {
	// this will execute the closure
	func() {
		fmt.Println(msg)
	}()
}

func invisible (msg string) {
	// gets run in the back ground
	go func() {
		fmt.Println(msg)
	} ()
}

func main() {
	ignored("Nothing to see")
	runnable("Drop the bomb")
	invisible("This is not visible")
}

