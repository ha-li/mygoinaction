package main

import (
	"math/rand"
	"time"
	"fmt"
)

func boring(msg string) {
	for i:=0; ;i++ {
		fmt.Println(msg, i)
		// time.Sleep(time.Second)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}


func example1 () {
	boring("Drop the bomb")
}

func example2() {
	// this will run boring in the background, like & in a shell command
	// so the program will not block and wait for the result
	// need channel for blocking
	go boring("Drop the bomb")
}


func example3() {
	// here we again run a go routine, so this will
	// run in the background, but since we go to sleep,
	// the process hangs around during the sleep, and
	// so the goroutine runs, and example3 runs
	// during the sleep, and so we see the output of the
	// goroutine
    go boring ("Dropping the bomb")
	fmt.Println("I'm listening")
	time.Sleep(1 * time.Minute)
	fmt.Println("You're boring, I'm Leaving")
}

func boring2 (msg string, c chan string) {
	for i:=0; ;i++ {
		// c<- blocks until a receiver is ready
		// here where are writing into the channel
		// the value of fmt.SPrintf, so
		// the sender is this go routine, sending
		// the results of the fmt.Sprintf into the
		// channel, and that does not complete
		// until a receiver gets the fmt.Sprint results
		c <- fmt.Sprintf ("%s %d", msg, i)

		// then we sleep for a bit, before we repeat it again forever
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func example4 () {
	//	this shows how channels work
	// channels connect a go routine
	// and the calling process
	// allowing them to communicate

	c:=make(chan string)

	// run the go routine, but pass it a
	// channel, so now the go routine and this function
	// are connected
	go boring2("Drop the bomb", c)

	// calling <-c will block until the
	// value is sent, the channel (c) is the sender
	// and the fmt.Printf is the receiver,
	// so nothing happens until the channel sends
	// a value, the channel has bbeen sent to the go routine
	// so the goroutine is the sender

	// next look at the go routine boring2
	for i:=0; i<5; i++ {

		// so as mentioned in the comments above
		// the for loop, <-c will block until
		// the value sent by the goroutine is received
		// here, so fmt.println is the receiver
		// of the fmt.Printf
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring, I'm Leaving")
}


// uncomment example 1, 2, 3, 4
// to see the progression
func main() {
	//example1()

	// example2()

	// example3()

	example4()

}
