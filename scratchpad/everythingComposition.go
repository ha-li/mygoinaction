package main

import (
	"fmt"
)
type Center struct {
	X float64
	Y float64
}

type Circle struct {
	Center
	Radius float64
}

func Resize (c Circle) {
	c.Radius = 28
	fmt.Println(c)
}


type BetterCircle struct {
	*Center
	Radius float64
}

func Shrink(c *BetterCircle) {
	c.Radius = 1
	fmt.Printf("%v %f\n", c.Center, c.Radius)
}

func main () {
	c := Circle {
		Center { X: 1.2, Y: 2.4},
		6.54,
	}
	fmt.Println(c)
	Resize (c)
	fmt.Println(c)

	b := BetterCircle {
		Center: &Center {X: 1.2, Y: 2.4},
		Radius: 6.54,
	}

	fmt.Printf("%v\n", b)
	Shrink(&b)
	fmt.Printf("%v\n", b)
}