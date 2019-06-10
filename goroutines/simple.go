package main

import (
	"fmt"
	"time"
)

func Announce(message string, delay time.Duration, c chan string) {
   time.Sleep(delay)
   c <- message
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	// var msg string
	go Announce("Hello World", 5000000, c1)
	go Announce("Done", 0, c2)
	go Announce("Hello 2 World", 500000000, c1)
	x := <- c1
	y := <- c2
	z := <- c1
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	//a := <- c1
	//fmt.Println(a)
}


