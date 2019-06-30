package main

import (
	"fmt"
)

func main () {
	var i1 []int

	// since i is un-initialized, it will be nil
	if i1 ==nil {
		fmt.Println("i1 is Nil")
	}

	// a nil value will still print out []
	fmt.Printf ("i1 is %v length of %d\n", i1, len(i1))

	// this uses the shorten syntax to declare and create
	// the variable, the type will be
	// determined by the initilizaed value
	i2 := []int {1, 2, 3}
	if i2 == nil {
		fmt.Println("i2 is Nil")
	}
	fmt.Println("i2 is ",i2)

	// this will not compile
	// because short form is used for both
	// declaration and initialization, in this case
	// there is no initialization, so it is a compile error
	//i3a := []int

	// i3b -- this is not a compile error since there is the {}
	// which will allocate a empty slice
	i3b := []int{}
	if i3b == nil {
		fmt.Println ("i3b is nil")
	}
	fmt.Printf("i3b is %v length of %d\n", i3b, len (i3b))



	// using make to allocate a slice requires you to pass
	// in the size of the slice
	// during the allocation, the slice will be initialized
	// to the equivalent nil value for that type
	// so i3 will be [0, 0, 0] after make
	i4 := make([]int, 3)
	if i4 == nil {
		fmt.Println( "i4 is nil")
	}
	i4[0] = 1
	fmt.Println("i4 is ", i4)

	// the rule of thumb when declaring variables is
	// 1. to initialize to their zero value, use var form
	// 2. if you want to add extra initialization or make a
	//    function call, use the short form
}
