package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "bob"
	fmt.Println(a[0])
	fmt.Println(a)

	b := [2]string{"bob", "keyll"}
	fmt.Println(b)

	c := [...]string{"one", "two", "three"}
	fmt.Println(c)

	fmt.Printf("%q\n", c)

	var d [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = fmt.Sprintf("row %d - col %d", i, j)

		}
	}

	fmt.Printf("%q\n", d)

	var i [10]int
	i[0] = 12
	i[1] = 10
	i[2] = 5
	for k, v := range i {
		fmt.Printf("%d %d\n", k, v)
	}
}
