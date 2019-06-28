package main

import (
	"fmt"
	//"time"
)

func doSelect(c1, c2 chan int) {
	for {
		select {
		case x:= <-c1:
			x = x*2
			fmt.Println(x)

		case y:= <-c2:
			y = y*3
			fmt.Println(y)
		default:
			break;
		}
	}
}

func doSomething (c chan int, n, o int) {
	for i:=0; i < n; i++ {
		c <- o
	}

}

func main() {
	c1 := make (chan int)
	c2 := make (chan int)
	//c3 := make (chan int)

	go doSomething (c1, 4, 1)
	go doSomething (c2, 2, 2)
	doSelect(c1, c2)
}