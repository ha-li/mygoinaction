package main

import (
	"fmt"
)

type Point struct {
	Center float64
	Radius float64
}

func mapLiterals() {
	// the syntax of a map declaration is
	// keyword 'map'
	// square bracket with the key type inside
	// outside is the value type
	// so this is a map of int that point to strings
	m := map[int]string {
		1 : "One",
		2 : "Two",
		3 : "Three",
	}

	for k,v := range m {
		fmt.Printf ("%d %s\n", k, v)
	}

	p := map[string]Point {
		"g": Point{2.4, 2.3},
		"h": Point{1.4, 5.3},
	}
	for s,t := range p {
		fmt.Printf ("%s %v\n", s, t)
	}
}

func mapMake () {
	// instead of a map literal you can also make to create a map
	m := make(map[int] string)
	m[1] = "One"
	m[2] = "Two"
	m[3] = "Three"

	for k,v := range m {
		fmt.Printf( "%d %s\n", k, v)
	}

	p := make(map[string]Point)
	p["g"] = Point{4.3, 43.}
	p["h"] = Point {67.3, 43.0}
	for s,t := range p {
		fmt.Printf ("%s %v\n", s, t)
	}
}

func mapFunctions () {
	//
	p := make(map[string]Point)

	// insert/update an element in a map
	p["g"] = Point{4.3, 43.}
	p["h"] = Point {67.3, 43.0}

	//delete an element in a map, here we delete the element of "g" and
	// "g" itself
	delete(p, "g")
	for s,t := range p {
		fmt.Printf ("%s %v\n", s, t)
	}

	// test if a key is in the map with a two value assignment
	// the 2nd value is a boolean
	v,ok := p["g"]
	if ok {
		fmt.Printf ("%v\n", v)
	} else {
		fmt.Println ("g was not found")
	}
}

func testingMap () {
	p := make(map[string]Point)

	// insert/update an element in a map
	p["g"] = Point{4.3, 43.}
	p["h"] = Point {67.3, 43.0}

	// when checking a map for a value, there are 2 return values
	// 1 - the actual return value
	// 2 - boolean indicating the key is found or not
	v, ok := p["a"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("a was not found")
	}

	v, ok = p["h"]
	if ok {
		fmt.Println("h was found ", v)
	}
	p["h"] = Point{5, 5}
	fmt.Println("after changing p[h], v is ", v)
	fmt.Println( "and p[h] is ", p)
	fmt.Println( "v is a copy of p[h] but not p[h]")
}

func main() {
	mapLiterals()
	mapMake()
	mapFunctions()
	testingMap()
}