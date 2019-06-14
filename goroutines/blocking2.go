package main

import (
	"fmt"
	"time"
)
// another example to show the blocking
// nature of channels

// we know channels are used for communicating between
// go routines and calling routines


func blocked (c chan string) {
	for {
		c <- "I am blocked"
	}
}

func main() {
	c := make (chan string)
	fmt.Println ("The go routine is blocked...")
	go blocked (c);

	time.Sleep(10 * time.Second)

	for i:=0; i < 10; i++ {
		fmt.Printf ("Unblocking whoever is blocked. %s\n", <-c)
	}
}
