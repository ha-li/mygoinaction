package main

import (
	"fmt"
)

func main() {
	celebs := map[string]int{
		"Bob Dylan":          50,
		"Taylor Swift":       25,
		"Jude Law":           43,
		"Scarlett Johansson": 30,
	}

	shapes := map[int]string{
		1: "dot",
		2: "line",
		3: "triangle",
		4: "square",
		5: "hexagon",
	}

	fmt.Printf("%v\n", celebs)
	fmt.Printf("%v", shapes)

	shapes[10] = "dexagon"
	fmt.Printf("%#v\n", shapes)
}
