package person

import (
	"fmt"
)

func (p Person) AltGreet(alt string) string {
	return fmt.Sprintf("%s %s %s", alt, p.FirstName, p.LastName)
}
