package main

import "fmt"

// returning a pointer to a int type
// returns the address space
func simple () *int {
	x := 3
	return &x
}


type Img struct {
	x int
}


func simplealso() *Img {
	i := Img {4}
	return &i
}


func main() {
	// gets the pointer/address
	y := simple()
	// printing the pointer gets you an address value
	fmt.Println(y)

	// dereferencing it will get you the value the pointer points at
	fmt.Println(*y)

	a := *simple()
	fmt.Printf("a = %d", string(a))

	z := simplealso()

	// these two are not the same
	fmt.Println(z)
	fmt.Println(*z)

	// these two are the same
	fmt.Println(z.x)
	fmt.Println((*z).x)

	// this does not compile
	// fmt.Println(*z.x)
}
