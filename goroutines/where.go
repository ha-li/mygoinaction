package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func where(name string, c chan int) {
	defer wg2.Done()

	// each get from a channel also returns true/false
	// if true, the channel is open, if false the channel was closed
	i, ok := <- c

	if ok {
		fmt.Printf("%s got %d\n", name, i)
		close (c)
		return
	} else {
		fmt.Printf ("%s did not get the ball\n", name)
	}
}

func main() {
	c := make (chan int)
	wg2.Add(2)

	go where ("Bob", c)
	go where ("Robert", c)
	c <- 0

	wg2.Wait()
	//close(c)
}
