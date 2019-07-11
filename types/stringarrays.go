package main

import (
	"fmt"
)

func main() {
	var sarray [5]string
	sarray[1] = "a"
	fmt.Printf("s array %v \n", sarray)

	var s2array [4]string
	s2array[0] = "bb"
	fmt.Printf("s array %v \n", s2array)

    s3array := [3]string {"c"}
	fmt.Printf("s array %v \n", s3array)

	s4array := [3]string {"dd"}
	fmt.Printf("s array %v \n", s4array)

	// this is equivalent to copying s3 into s4 (ie with a for loop)
	// but is shorter; if you change s3, s4 does not change
 	s4array = s3array   // s4 is gone, s3 and s4 have the same value
	fmt.Printf("s3array %v \n", s3array)
	fmt.Printf("s4array %v \n", s4array)


	s4array[0] = "a new string"  // only changing s4, not s3
	fmt.Printf("s3array %v \n", s3array)
	fmt.Printf("s4array %v \n", s4array)

	// declare string pointer array
    var sp1 [3]*string
	sp2 := [3]*string{new(string), new(string), new(string)}
	fmt.Printf("sp1 %v \n", sp1)
	fmt.Printf("sp2 %v \n", sp2)

	*sp2[1] = "Green"
	fmt.Printf("sp1 %v \n", sp1)
	fmt.Printf("sp2 %v \n", sp2)

	// now sp1 and sp3 point to the same memory spaces
	// so if you change one, you will change the other
	sp1 = sp2
	fmt.Printf("sp1 %v \n", sp1)
	fmt.Printf("sp2 %v \n", sp2)

	*sp1[0] = "Purple"
	fmt.Printf("sp1 %v \n", sp1)
	fmt.Printf("sp2 %v \n", sp2)

}
