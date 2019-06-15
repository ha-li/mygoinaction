package main

import (
	"fmt"
	"time"
)

type Message struct {
	str string
	wait chan bool
}


func method (c chan string, n int) {
	for {
		c <- fmt.Sprintf("Method %d", n)
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	c1 := make (chan string)
	c2 := make (chan string)
	c3 := make (chan int)

	//var v1, v2 string

	go method (c1, 1)
	go method (c2, 2)

	for {
		select {
		case v1 := <-c1:
			fmt.Printf("received %v from c1\n", v1)
		case v2 := <-c2:
			fmt.Printf("received %v from c2\n", v2)
		case c3 <- 23:
			fmt.Printf("sent %v to c3\n", 23)
		default:
			fmt.Printf("no one was ready\n")
		}
		time.Sleep(1 * time.Second)
	}
}