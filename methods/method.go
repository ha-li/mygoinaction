package main

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName string
}

// this method belongs to struct User
// as denoted by (u User)
func (u User) Greet(s string) string {
	return fmt.Sprintf("%s %s says '%s'\n", u.FirstName, u.LastName, s)
}

func main() {
	a := User{"Bob", "Hanigan"}
	fmt.Println(a.Greet("Bye!"))
}
