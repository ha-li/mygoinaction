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

func main () {
	c := fanIn(boring5("Joe"), boring5("Ann"))
	for i:=0; i < 10; i++ {
		fmt.Print(<-c)
	}

	fmt.Println("You're both boring")
}