package main

import (
	"fmt"
	"sync"
)



func main () {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	s := make (chan string)
	go func (msg string, c chan string) {
		for {
			c <- msg
			waitGroup.Done()
		}
		//waitGroup.Done()
	}("Try to Stop Me", s)

	//waitGroup.Wait()
	fmt.Println ("outputting ", <- s)
	fmt.Println("All Done")
}