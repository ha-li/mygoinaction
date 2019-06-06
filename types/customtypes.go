package main

import (
	"fmt"
)

type MyOwnType struct {
	Name string
	Age  int
	BirthDate string
}


func main() {
	var one = MyOwnType {
		"bob", 12, "never",
	}

	fmt.Println (one)

	var two = MyOwnType {
		Age: 12, Name: "Bob",
	}

	fmt.Println(two)

	fmt.Println (two.Name)
}