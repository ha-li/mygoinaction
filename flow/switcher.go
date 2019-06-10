package main

import (
	"fmt"
	"time"
)

func main() {
	score := 7

	switch score {
	case 0, 1, 2:
		fmt.Println("Bad score")
	case 3, 4, 5:
		fmt.Println("Not Bad")
	case 6, 7:
		fmt.Println("Good")
	default:
		fmt.Println("Amazing")
	}

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Unix()) // time as long epoch
	fmt.Println(t.Unix() % 2)
}
