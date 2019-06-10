package main

import "fmt"

func main() {
	a := [...]string{"hello", "world"}
	fmt.Printf("%q\n", a)

	fmt.Printf("%s\n", a)

	var b [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = fmt.Sprintf("row %d - column %d", i, j)
		}
	}

	fmt.Printf("%q\n", b)

	//var c [2]string
	// this will not compile
	//c[3] = "bob"
	//fmt.Printf(c)

	// 'range' is a for loop operator that
	// returns the index and value of a loop
	var i = []int{1, 3, 2, 8, 5}
	for j, k := range i {
		fmt.Printf("%d, %d\n", j, k)
	}

	// use _ to ignore a value
	for _, y := range i {
		fmt.Println(y)
	}

	for x, _ := range i {
		fmt.Println(x)
	}

	for x := range i {
		fmt.Println(x)
	}

	var s = []string{"me", "myself", "i"}
	for _, sn := range s {
		fmt.Printf("%s\n", sn)
	}
}
