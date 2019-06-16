package main

import (
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Location string
}

func (u *User) Greet() {

	fmt.Println ("Hello ", u.Name)
}

// player is composed of User also
// a player will also inherit the users functions
// so player will have  a Greet method, see below
type Player struct {
	User
	Number int
	GameId string
}

func main() {
	u := User{12, "Jason", "USA"}

	fmt.Println(u)

	p := Player{u, 12, "Football"}
	fmt.Println(p)

	p2 := Player{}
	p2.Id = 12
	p2.Name = "Jim"
	p2.Location = "USA"
	p2.Number = 125
	p2.GameId = "Tether"
	fmt.Println(p2)

	p3 := Player{
		User{Id: 19, Name: "Drogo", Location: "Las Vegas"},
		94021,
		"Footer",
	}
	fmt.Println(p3)

	// methods are inherited in a composition
	// relationship, which allows for very powerful
	// object building, for example, if an Job struct has
	// a Logger struct, we can do things like j.log(...)
	// an as a result, Job will also implement
	// any interface that Logger also implements, useful
	// for certain 'polymorphic' behaviors, ie polymorphism
	// is automatically done if you have a composition
	p3.Greet()
}
