package main

import (
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Location string
}

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
}
