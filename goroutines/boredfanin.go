package main

import (
	"math/rand"
	"time"
	"fmt"
)

func boring5(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i:=0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3))* time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make (chan string)
	go func() {
		for {
			c <- <- input1
		}
	}()

	go func() {
		for {
			c <- <- input2
		}
	}()
	return c
}

func fanIn2(input1, input2 <-chan string) <-chan string {
	c := make (chan string)
	go func() {
		for {
			select {
			case s := <-input1: c <- s
			case s := <-input2: c <- s
			}
		}
	}()
	return c
}

func main () {
	// two differents ways implementations that illustrate how select works
	// in fanIn, we create 2 channels, infinite loop and call
	// go routines on each of them, which in turn each block until msg come in
	//
	// in fanIn2, we create 2 channels, infinitely loop,
	// and select on the two inputs, and when a channel is
	// ready, it gets selected

	// comment out fanIn or fanIn2 to show 2 implementations of the same thing
	//c := fanIn(boring5("Joe"), boring5("Ann"))
	c := fanIn2(boring5("Joe"), boring5("Ann"))
	for i:=0; i < 10; i++ {
		fmt.Print(<-c)
	}

	fmt.Println("You're both boring")
}