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

	fmt.Printf("%v %d \n", celebs, len(celebs))
	celebs["Coffee Shop Barista"] = 28
	fmt.Printf("%v %d \n", celebs, len(celebs))

	fmt.Printf("%v %d \n", shapes, len(shapes))

	shapes[10] = "dexagon"
	fmt.Printf("%#v\n", shapes)

	// this is a nil map (a declaration without any initialization)
	// this nil map can never hold any values
	var nilmap map[int]string
	// a nil map looks like an empty map, has a length of zero
	fmt.Printf("nilmap %v %d \n", nilmap, len(nilmap))

	// but as soon as you try to populate it, it will cause a run time panic
	nilmap[3] = "three"
	//fmt.Printf("%v %d \n", nilmap, len(nilmap))

	actor := "Jude Law"
	v,exst := celebs[actor]
	if exst {
		fmt.Printf ("%s is %d years old\n", actor, v)
	} else {
		fmt.Printf ("%s is not in the map \n", actor)
	}
    remove (celebs, actor)
	_,exst = celebs[actor]
	if exst {
		fmt.Printf ("%s is still in the map\n", actor)
	} else {
		fmt.Printf ("%s has been removed from the map\n", actor)
	}

}

// maps are like slices in that they encapsulate a data structure
// and so even though they are passed by value, changes inside
// the function are permanent
func remove (names map[string]int, name string) {
   delete (names, name)
}