package main

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
}

// this method belongs to struct User
// as denoted by (u User)
func (u User) Greet(s string) string {
	return fmt.Sprintf("%s %s says '%s'\n", u.FirstName, u.LastName, s)
}

func local(n string) (region, continent string) {
	switch n {
	case "New York", "LA", "Chicago":
		continent = "North America"
	default:
		continent = "Unknown"
	}

	return
}

func main() {
	a := User{"Bob", "Hanigan"}
	fmt.Println(a.Greet("Bye!"))

	test1, test2 := local("Santa Monica")
	fmt.Printf("test1: %s, test2: %s", test1, test2)
}
