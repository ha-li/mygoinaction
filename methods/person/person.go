package person

import (
	"fmt"
)

type Person struct {
	FirstName, LastName string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello %s %s", p.FirstName, p.LastName)
}
