package main

import (
	"fmt"
)

func main () {
	// a slice of lenght 3, capacity 3
	slice := []int {10, 20, 30}
	fmt.Println(slice)

	// a slice of length 100, capacity 100
	slice = []int { 99: 0 }
	fmt.Println(slice)

	var nilslice []int
	fmt.Printf("nil slice %v, len %d, cap %d\n", nilslice, len(nilslice), cap(nilslice))
	if nilslice == nil {
		fmt.Println ("nil slice is nil")
	} else {
		fmt.Println ("nil slice is not nil")
	}

	slice2 := []int {}
	fmt.Printf("empty slice %v, len %d, cap %d\n", slice2, len(slice2), cap(slice2))
	if slice2 == nil {
		fmt.Println ("slice2 is nil")
	} else {
		fmt.Println ("slice2 is not nil")
	}

	newslice := slice[2:3]
	// the starting value of newslice is slice[2], and the last
	// value of newslice is not inclusive of slice[3]
	// so newslice = slice[2,3), which gives newslice a
	// length of 3 - 2
	newslice[0] = 1
	fmt.Printf("new slice %v, len %d, cap %d \n", newslice, len(newslice), cap(newslice))
	fmt.Printf("original slice is also changed where newslice is changed %v\n", slice)

    newslice = append(newslice, 80)
	fmt.Printf("new slice %v, len %d, cap %d \n", newslice, len(newslice), cap(newslice))
	fmt.Printf("original slice is also changed where newslice is changed %v\n", slice)


    slice = []int {4: 1}
	fmt.Printf("changing original slice %v\n", slice)

    // take a new slice that is of length 1 which is the last value of slice, so 1
	newslice = slice[4:5]
	fmt.Printf("new slice %v, len %d, cap %d \n", newslice, len(newslice), cap(newslice))
	fmt.Printf("original slice %v\n", slice)

	// append 800 to new slice, so new slice is now 1, 800
	// but that is larger then original slice, so what happens to original slice?
	// it should stay the same, because Go will allocate a new slice
	// and return it from append, so now newslice and slice are different
	newslice = append(newslice, 800)
	fmt.Printf("new slice %v, len %d, cap %d \n", newslice, len(newslice), cap(newslice))
	fmt.Printf("original slice %v\n", slice)


    s := []string {"one", "two", "three", "four", "five"}
    for i,v := range s {
    	fmt.Printf("copied v %v, trying to change to Dumb\n", v)
    	v = "Dumb"
    	fmt.Printf("original value %v, changing by Index\n", s[i])
    	s[i] = "Better"
		fmt.Printf("new value %v\n", s[i])
	}

    fmt.Printf("%v\n", s)

    test(s)

    for _, v := range s {
    	fmt.Printf("this should be 'New' %v \n", v)
	}
}

func test(slice []string) {
	for i, _ := range slice {
		slice[i] = "New"
	}
}
