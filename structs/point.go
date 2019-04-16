package main

import "fmt"

type Point struct {
	X int
	Y int
}

var (
	p = Point{X: 1, Y:2}
	c = &Point{Y: 4, X: 1}
	d = Point{X:1}
	f = Point{}
)

func main () {
	fmt.Println(p)
	fmt.Println(*c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println((*c).X)
}
