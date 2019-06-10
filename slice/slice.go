package main

import (
	"fmt"
)

func main() {
	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)
	fmt.Println(b[0:0])

	var z []int
	if z == nil {
		fmt.Println("nil")
	}

	// range for iterating over a data structure like array/slice/map
	for i, v := range b {
		fmt.Printf("%d %d\n", i, v)
	}
}
