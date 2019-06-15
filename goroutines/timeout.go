package main

import (
	"math/rand"
	"time"
	"fmt"
)

// time.After is a method that returns a channel that
// blocks for the specified time
// after the interval, the channel delivers the current time one
func boring6(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i:=0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3))* time.Millisecond)
		}
	}()
	return c
}

// we try to timeout after each message,
// but never get a time because the for loop is too fast
// so s:= <- will always send before 1 second
// each for loop iteration resets again
func example10() {
	c := boring6("Joe")
	for {
		select {
		case s := <-c:
		   fmt.Println(s)
		case <-time.After(1*time.Second):
		   fmt.Println("You're too slow")
		   return
		}
	}
}

// here we set the timer once before the conversation
// starts, so the for loop will terminate
func example20(){
	c := boring6("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		   case s := <-c:
		   	   fmt.Println(s)
		   case <-timeout:
		   	   fmt.Println("You talk too much.")
		   	   return
		}
	}
}

func main() {
	//example10()
	example20()
}
