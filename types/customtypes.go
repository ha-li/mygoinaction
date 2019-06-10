//usr/local/go/bin/go run "$0" "$@"; exit "$?"
package main

import (
	"fmt"
)

type MyOwnType struct {
	Name      string
	Age       int
	BirthDate string
}

func NewMyOwnType(name string, age int, bday string) *MyOwnType {
	// new expression allocates a zeroed value of the requested type
	// and returns a pointer to it
	m := new(MyOwnType)
	m.Name = name
	m.Age = age
	m.BirthDate = bday
	return m
}

func MakeMyOwnType(name string, age int, bday string) MyOwnType {
	m := NewMyOwnType(name, age, bday)
	return *m
}

func main() {
	var one = MyOwnType{
		"bob", 12, "never",
	}

	fmt.Println(one)

	var two = MyOwnType{
		Age: 12, Name: "Bob",
	}

	fmt.Println(two)

	fmt.Println(two.Name)

	var n = NewMyOwnType("Bob", 112, "Older")
	fmt.Println(n)

	o := MakeMyOwnType("Randy", 32, "younger")
	fmt.Println(o)
}
