package main

import (
	"fmt"
)

func main() {
	var bob = "bob"
	fmt.Println(bob)

	bob2 := "bob2"
    fmt.Println(bob2)

	age := 43
	fmt.Println(age)

	var age2 int = 12
	fmt.Println(age2)

	action := func (n string) {
		fmt.Print(n)
		fmt.Println (age2)
	}

	action("bob is ")
}
