package main


import (
	"fmt"
)

// the builtin function 'make' can be used on
// slices,
// maps,
// channels
// but not arrays
func makeArrayInt(){
	fmt.Println("You cannot 'make' an array in Go, instead, see makeSliceX")
}

func makeSliceInt() {
	a := make([]int, 3)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Printf ("%v\n", a)
}

func makeMapInt () {
	// the 2 is not relevant, so it is ignored
	m := make(map[int]string, 2)
    m[0] = "Bobby"
    m[1] = "Tether"
    m[2] = "Holy Cow"
    m[3] = "Money"
    fmt.Printf ("%v\n", m)
}

func makeChannel () chan int {
	c := make(chan int)
	return c
}

func main() {
	makeArrayInt()
	makeSliceInt()
	makeMapInt()


	c := makeChannel()
	fmt.Println(c)
}