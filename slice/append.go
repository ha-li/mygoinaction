package main

import "fmt"

func main() {
	v := []string {"a", "bb", "c"}
	w := []string {"d", "e", "f"}

	v = append(v, w...)
	fmt.Println(v);
}
