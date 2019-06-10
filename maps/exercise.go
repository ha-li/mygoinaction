package main

import "fmt"

var names = []string{"Katrina", "Even", "Raymond", "Adam"}

func maxLength(names []string) int {
	maxLength := 0
	for _, name := range names {
		l := len(name)
		if l > maxLength {
			maxLength = l
		}
	}
	return maxLength
}

//sort each into a slice base on its length
func main() {

	var maxLength = maxLength(names)
	//fmt.Println (maxLength)
	// allocated a slice to hold the names by length
	sorted := make([][]string, maxLength)

	for _, name := range names {
		size := len(name)
		fmt.Printf("%s, %d\n", name, size)
		sorted[size-1] = append(sorted[size-1], name)
	}

	fmt.Println(sorted)
}
