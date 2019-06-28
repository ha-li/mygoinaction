package main

import (
	"fmt"
)

func echo (c chan string, msg string) {
	c <- msg
}

func main () {
	// this is the syntax for sending something into a channel
	// ch <- v

	// this is the syntax for receiving something from a channel
	// notice that the syntax says v is assigned the value
	// that is received from the channel (:= means use the type dynamically)
	// v := <- ch

    // by default channels block until the other side is ready,
    // whether that is reciving or sending a message
    // allowing you to synchronize without locks or condition variables

    // these three lines do not work at run time
    // channels must be used in conjunction with go routines.
    // channels block until the both the sender and receiver are ready
    // in this cause the send command blocks until the receiver is ready
    // which prevents the next line of code from executing

    // since the next line cannot progress, the program stops executing and
    // never progresses, which is why the compiler gives a deadlock warning

    //b := make (chan string)
    //b <- "this will not run, causes deadlock"
    //fmt.Println(<- b)

    // buffered channels can be used without go routines
    // they will not block block until the buffer is full
    // so this channel allows 1 until of write before blocking
    // so this is ok, then the next line clears the buffer so all is good
    c := make(chan string, 1)
    c <- "test"
    fmt.Println (<-c)

    // the correct way to use channels is with go routines
    // because the go routine executes in a separate thread
    // that continues to run regardless of this thread
    // so when you write into a channel in a go routine
    // that blocks, but the main thread can continue to execute
    // and if you are clearing the channel on the main thread,
    // this unblocks the go routine thread allowing it to continue
    // to run while the main thread is not blocked so it runs
    // until completion
    s := make (chan string)
    go echo (s, "test2")
    fmt.Println(<-s)


}
