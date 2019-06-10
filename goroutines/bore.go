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

func boring2 (msg string, c chan string) {
   for i:=0; ;i++ {
       c <- fmt.Sprintf ("%s %d", msg, i)
	   time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
   }
}

func main() {
	// this first part illustrates goroutines are like running shell commands with &
	go boring("Drop the bomb")
	fmt.Println("I'm listening")
	time.Sleep(1 * time.Minute)
	fmt.Println("You're boring, I'm Leaving")

//	this shows how channels work
	c:=make(chan string)
	go boring2("Drop the bomb", c)
	for i:=0; i<5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring, I'm Leaving")

}
