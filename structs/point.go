package main

import "fmt"

type Point struct {
	X int
	Y int
}

var (
	p = Point{X: 1, Y: 2}
	c = &Point{Y: 4, X: 1}

	// exmaples of named variable syntax,
	// where only initialize a subset of the fields of Point
	d = Point{X: 1}
	f = Point{}
)

func main() {
	fmt.Println(p)
	fmt.Println(*c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println((*c).X)

	// the new construct returns a pointer to the
	// zeroed struct, is equivalent to g:=&Point
	g := new(Point)
	fmt.Println("g:", g)

	h := &Point{}
	fmt.Println("h:", h)

	// this is true
	fmt.Println( "*h == *g", *h == *g)

	// this is false because it is testing the address value is equivalent?
	// which is not the case
	fmt.Println ("h == g", h == g)
}
