package main

import (
	"fmt"
	"github.com/mygoinaction/methods/person"
)

func main() {
	p := person.Person{"Bob", "Hope"}
	fmt.Println(p.Greet())
	fmt.Println(p.AltGreet("Bye"))
}
