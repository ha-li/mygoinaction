package main

import ("fmt"
	"math/rand"
	"time"
)

// we can use a channel as a service provider
func boring4(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i:=0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3))* time.Millisecond)
		}
	}()
	return c
}

func main() {
	joe := boring4("Joe")
	ann := boring4("Ann")

	for i:=0; i<5; i++ {
		fmt.Println (<-joe)
		fmt.Println(<-ann)
	}
}