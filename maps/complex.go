package main

import (
	"fmt"
)

type Vertex struct {
	Lat float64
	Long float64
}

var v map[string]Vertex

func main() {
	v = make(map[string]Vertex)
	v["Google Headquarters"] = Vertex{3.234, 2.23}
	fmt.Printf("%v", v);
}