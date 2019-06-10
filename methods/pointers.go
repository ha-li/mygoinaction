package main

import (
	"fmt"
)


type Vertex struct {
	X, Y float64
}

func (v Vertex) Scale(s float64) {
	v.X = v.X * s
	v.Y = v.Y * s
}

func (v *Vertex) BetterScale(s float64) {
	v.X = v.X * s
	v.Y = v.Y * s
}


func main() {
	v := Vertex {3.45, 643.2}
    fmt.Println(v)
	v.Scale(2)
	fmt.Println(v)
	v.BetterScale(2)
	fmt.Println(v)
}
