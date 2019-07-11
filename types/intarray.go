package main

import (
	"fmt"
)

func main() {
	var i0 [2]int
	i1 := [2]int {1, 2}
	fmt.Printf ("i0 %v\n", i0)
	fmt.Printf ("i1 %v\n", i1)

	// copy i1 over into i0? or do they point to the same values
	i0 = i1
	fmt.Printf ("i0 %v\n", i0)
	fmt.Printf ("i1 %v\n", i1)

	// this shows that it was a copied over,
	// since changing it will only change i0
	i0[0] = 2
	fmt.Printf ("i0 %v\n", i0)
	fmt.Printf ("i1 %v\n", i1)

}