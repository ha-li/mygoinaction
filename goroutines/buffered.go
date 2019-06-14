package main

import (
	"fmt"
	"time"
)

func send(c chan int, num int) {
	for i:=0; i < num; i++ {
		c <- num
	}
}


func receive (c chan int, num int) {
	for i:=0; i<num; i++ {
		fmt.Println (<- c)
	}
}

func main () {
	c := make (chan int, 5)

	go send (c, 10)


	// this will not cause deadlock
	// because go routines will not block the main thread.
	// the go routine will wait until the channel is available,
	// which happens in the main thread, so the go routine
	// is waiting for the main thread to clear the channel
	time.Sleep(10 * time.Second)
	for i:=0;i<10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("Running 2nd test")

	go send(c, 10)
	go receive (c, 10)
	time.Sleep(10*time.Second)
}