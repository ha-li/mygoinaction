package main

import "github.com/mygoinaction/methods/person"
import "fmt"

type Alias person.Person

func (a Alias) Ninja() string {
	return fmt.Sprintf("%s %s is a ninja", a.FirstName, a.LastName)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		f = -f
	}
	return float64(f)
}

func main() {
	a := Alias{"Bob", "Baker"}
	fmt.Println(a.Ninja())

	f := MyFloat(-0.54)
	fmt.Println(f.Abs())
}
