package main

import (
	"fmt"
)

type Maze struct {
    Wall int
}

func (m Maze) Type () {
	fmt.Printf("%T\n", m)
}

func main() {
	// a nil array
	var n []int

	// but when we print it out, it will show []
	fmt.Println(n)
	// even though it evaluates to nil
	if n != nil {
		fmt.Println(n)
	} else {
		fmt.Println("nil")
	}

	var c bool
	fmt.Println(c)

	// an empty string prints nothing
	var s string
	fmt.Println(s)

	// an uninitialized map
	var m map[string]string
	// printing it out will show map[]
	fmt.Println(m)
	// but evaluating it is nil
	if m != nil {
		fmt.Println(m)
	}else {
		fmt.Println ("m map is nil")
	}

	var maze Maze
	if m != nil {
		fmt.Println(maze)
	}else {
		fmt.Println(" maze is nil")
	}
	// can still call method on a nil object
	maze.Type()
}