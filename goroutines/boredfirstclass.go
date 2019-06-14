package main

import ("time"; "fmt"; "math/rand")


// another example that channels are first class entities
func boring3(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i:=0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3))* time.Millisecond)
		}
	}()
	return c
}

// channels can be passed into functions and returned from
// functions, making them a first class entity
func main() {
	c := boring3("boring")
	for i:=0; i<5; i++ {
		fmt.Printf("You say: %q \n", <-c)
	}
	fmt.Println("You're boring, I'm leaving")
}