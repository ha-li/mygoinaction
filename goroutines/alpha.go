package main

import (
	"fmt"
	"sync"
)

func main () {

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting the alphabet...")

	// create a go routine
	go func () {
		defer wg.Done()
		for n := 0; n < 3; n++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char);
			}
		}
	}()

	// create another go routine
	go func () {
		defer wg.Done()
		for n := 0; n < 3; n++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// once the go routines are created, this will run immediately
	// so this is likely to run before the go routines are run
	fmt.Println("finishing...")

	// tell the main process to wait for the 2 go routines
	wg.Wait()


	// this is the last to run
	fmt.Println ("Waited for everything")
}
