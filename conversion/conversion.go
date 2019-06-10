package main

import (
	"fmt"
)

var (
	i int     = 42
	f float64 = float64(i)
	u uint    = uint(f)
)

func main() {
	//f = "%s\n"
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)
}
