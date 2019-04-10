package main


import (
	"fmt"
)

var (
	name, locations string
	age  int
	//location string
)

var (
	province string = "BC"
	teamCode int = 5
)

var (
	sex string = "m"
)

const (
	OK = 200
	Pi = 3.14
)

func main () {
	name = "Hao"
	age = 24
	locations = "California"
	talent := "juggling"
	fmt.Println (name, age, location)
	fmt.Println (teamCode, province)
	fmt.Println (talent)
	fmt.Println (Pi)
	fmt.Println (OK)

	cylonModel := 5
	fmt.Println (cylonModel)

	cyType := "Boo"
	fmt.Println (cyType)

	name := "Caprix-Six"
	aka := fmt.Sprintf("Number %d", 6)
	fmt.Printf("%s is also known as %s", name, aka)
	fmt.Println()
	//print ("Hello World")
	fmt.Println (add(3,4))
	fmt.Println (sub(4, 3))

	state, continent := location("LA");

	fmt.Println (state)
	fmt.Println (continent)

	var hometown, con = location("New York")
	fmt.Println (hometown)
	fmt.Println (con)

	region, continent := whereIlive("Matt", "New York")
	fmt.Printf ("%s lives in %s", region, continent)
}

func whereIlive (name, city string) (region, continent string) {
	switch city {
	case "New York":
		continent = "North America"
	}

	return
}

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York":
		region, continent = "New York", "North America"
	}

	return region, continent
}

func add (x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

