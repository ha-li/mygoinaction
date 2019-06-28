package main

import (
	"fmt"
	"time"
)

// range is an operator of go that can be used for
// arrays/slices
// maps
// strings

// then we show a fun example that would be a good interview question


// a simple example showing the for loop syntax and
// the range builtin that returns 2 values, the index and the value
// (on a array/slice)
func forRange() {
	c := make([]int, 20)

	// the first return value is the index, the 2nd return value is the value of the index
	// lets see if this is congruent with other uses of range
	for i, v := range(c) {
		fmt.Printf("index %d value %v\n", i, v)
	}
}

func mapRange() {
	m := map[int]string {
		0: "Zero",
		1: "One",
		2: "Two",
	}

	// the first returnvalue is the key, the 2nd return value is the value
	// of the key in the map, like the range operator on a slice/array
	// the 2nd value is the value, so that is kinda consistent
	for k, v := range(m) {
		fmt.Printf ("%d %s\n", k, v)
	}
}

func stringRange() {
	s:= "simplestring"

	// range can be used on a string, why (its just a char array)
	// in which case, in consistency with the other range,
	// expect the 1st return value to be an index, the 2nd value to be
	// the value at that index
	for i, c := range(s) {
		// need to convert the c into a string
		fmt.Printf ("%d %s\n", i, string(c))
	}
}

func funRange () {
	a := []int {1, 2, 3}
	// does this terminate?
	// yes, because range is taken before the start of the loop
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
		a = append(a, i)
	}
}

func goRoutine (n int, c chan string) {

	for i := 0; i < n; i++ {
		c <- fmt.Sprintf("%s %d\n", "hello", i)
		//time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Second * 10)
	close(c)
}

func channelRange() {
	c := make(chan string)
    go goRoutine(10, c)

	// range (channel) requires a channel be closed
	// this will for loop until the channel is
	// closed and we may not see all the "hello"s
	for range(c) {
		fmt.Println(<-c)
	}

	// there we are quaranteed to see all the "hello"
	/* for i:=0; i < 10; i++ {
		fmt.Println (<-c)
	} */
}

func main() {
	forRange()
	mapRange()
	stringRange()
	funRange()
	channelRange()
}