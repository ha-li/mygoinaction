package main

import (
	"fmt"
)

type Student struct {
	Email string
	Name string
}

// this will not work because a copy of s is made during the call
func (s Student) changeEmail (email string) {
	s.Email = email
}

// this will work because s is shared with the method
func (s *Student) changeName (name string) {
	s.Name = name
}


func main() {
	s1 := Student {
		"bob@yahoo.com",
		"Bob Iger",
	}
	fmt.Println(s1)
	fmt.Println ("Changing s1 emails")
	s1.changeEmail ("bob.iger@gmail.com")
	fmt.Println(s1)


	s2 := Student {"tesla@gmail.com", "Elon Musk"}
	fmt.Println(s2)
	s2.changeName("Mary Mayer")
    fmt.Println(s2)
}