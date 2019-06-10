package main

import "fmt"


type Person struct {
	LastName, FirstName string
}

func (p Person) Name() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

type Pet struct {
	MyName string
}

func (p Pet) Name() string {
	return fmt.Sprintf("%s", p.MyName)
}

type Name interface {
	Name() string
}

func Greet(n Name) string {
	return fmt.Sprintf("hello %s", n.Name())
}

func main() {
	p := Pet {"Spike"}
	p2 := Person {"Bob", "Candida"}
	fmt.Println(Greet(p))
	fmt.Println(Greet(p2))
}