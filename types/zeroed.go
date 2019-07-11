package main

import (
	"fmt"
)

func main () {
	var s string
	fmt.Printf("a nil string is %v \n", s)

	var i int
	fmt.Printf(" a nil int is %d\n", i)

	var b bool
	fmt.Printf("a nil boolean is %v \n", b)

	var sarray [3]string
	fmt.Printf("a nil string array is %v, %d\n", sarray, len(sarray))

	var iarray [3]int
	fmt.Printf("a nil int array is %v, %d\n", iarray, len(iarray))

	i2array := [3]int {1, 2, 3}
	fmt.Printf("an initialized array is %v\n", i2array)

	i3array := [...]int {1, 2, 3, 4, 5}
	fmt.Printf("an initialized array is %v\n", i3array)

	var parray [3]*int
	fmt.Printf("a nil *int array is %v\n", parray)

	p2array := [3]*int{0: new(int), 1:new(int), 2:new(int)}
	*p2array[0] = 5
	fmt.Printf("a partially initialized *int array %v\n", p2array)
	fmt.Printf("a index %d (the address space???) \n", p2array[0])
	fmt.Printf("a index %d (actual value of) \n", *p2array[0])

	var arrays1 [4]string
	arrays2 := [4]string {"Zero", "One", "Two", "Three"}
	fmt.Printf("arrays1: %v\n", arrays1)
	fmt.Printf("arrays2: %v\n", arrays2)

	// these next few lines show the effect of Go's copy by value
	// arrays1 is assigned to the same value as arrays2, so
	// copy by value means arrays1 gets a copy of whatever arrays2 is,
	// then we change s2[1], and printing out the values shows
	// that s2 and s1 are not the same anymore, so they are actually
	// independent arrays
	arrays1 = arrays2
	arrays2[1] = "Changed One"
	fmt.Printf("arrays1: %v\n", arrays1)
	fmt.Printf("arrays2: %v\n", arrays2)

	// these next few lines show the same effect with int arrays
    var i1 [4]int
	i2 := [4]int {1, 2, 3, 4}
	fmt.Printf("i1: %v\n", i1)
	fmt.Printf("i2: %v\n", i2)

	i1 = i2
	fmt.Printf("i1: %v\n", i1)
	fmt.Printf("i2: %v\n", i2)

	i2[1] = 25
	fmt.Printf("i1: %v\n", i1)
	fmt.Printf("i2: %v\n", i2)

	// to declare a pointer to an array of size 4
	// but remember that the preferred way to work with pointers to arrays
	// in go is to use a slice
	var i3 *[4]int
	i4 := [...]int {1, 2, 3, 4}
	fmt.Printf("i3: %v\n", i3)
	fmt.Printf("i4: %v\n", i4)

	i3 = &i4
	fmt.Printf("i3: %v\n", *i3)
	fmt.Printf("i3: %v\n", i3)
	fmt.Printf("i4: %v\n", i4)

	var i5 []int

	print(i5)

	var ii [2]int
	iii := [2]int { 1, 2}
	fmt.Printf("%v %v\n", ii, iii)
	ii = iii
	fmt.Printf("%v %v\n", ii, iii)

	iv := [3]int { 1, 2, 3}
	fmt.Printf("%v %v\n", iv, iii)

	// a compiler error because of wrong type
	// with arrays, the dimension (size) of the array is part of its type
	iii = iv
	fmt.Printf("%v %v\n", iv, iii)
}

// a function that takes a slice of int
func print(array []int ) {
	fmt.Printf("array %v\n", array)
}