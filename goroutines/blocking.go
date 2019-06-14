package main

import (
	"fmt"
	"time"
)

func sender (c chan string, num int, msg string) {

	time.Sleep(5 * time.Second)
	for i := 0; i < num; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
	}
}


func main() {
	c := make (chan string)
	n := 50

	// here we call a go routine with
	// a channel so this main is connected
	// with the go routine
	// the go routine for loops 50 times,
	// each time writing to the channel
	// so that will block until a receiver
	// gets the message
	go sender (c, n, "Booyah")

	// here we have a receiver (fmt) to the channel,
	// that iterates 50  times,
	// so we will get each of the message
	for i:=0;i<n;i++ {
		fmt.Printf("%s\n", <-c)
	}

	// we do the same thing again, so again
	// the sender will block (stop executing)
	// until a receiver gets the message
	go sender (c, n, "Will i See?")
	// but this time, we don't have a receiver
	// we just println "Not gonna wait" instead of receiving the
	// value, so we will not see "Will i See"
	for i:=0;i<n; i++ {
		fmt.Println("Not gonna wait")
	}
    // after 50 iterations of Not gonna wait, we will end
    // and just ignore "Not gonna see", because
    // the block stops the go routine from executing
	fmt.Println("All Done")
}
