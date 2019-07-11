package main
import (
	"fmt"
	"time"
)

// this will copy the array so it will be slow and consumes alot of memory
func slow(array [1000000]int) {
	fmt.Println("a quick print")
}

// this will copy the address, so it will be fast and not consumer much memory
func fast(array *[1000000]int) {
	fmt.Println("a quick print")
}

func main() {
	array := [1000000]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	start := time.Now()
	slow(array)
	slow := time.Since(start)
	fmt.Println("total time ", slow)

	start = time.Now()
	fast(&array)
	fast := time.Since(start)
	fmt.Println("total time ", fast)

	//fmt.Printf("fast is %f times faster than slow \n", (fast/slow))
}


