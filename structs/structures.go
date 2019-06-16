package main

import (
	"fmt"
	"time"
)


// a struct is a collection of fields/properties
// structs do not have methods, but you can crate a method
// that belongs to a struct
type Bootcamp struct {
	MyLat  float64
	MyLon  float64
	MyDate time.Time
}

func (b Bootcamp) getInfo () {

	fmt.Printf ("Location: Lat(%f) Long(%f),\nDate: %v\n", b.MyLat, b.MyLon, b.MyDate)
	fmt.Printf ("Location: Lat(%f) Long(%f),\nDate: %s\n", b.MyLat, b.MyLon, b.MyDate)

}

func main() {
	var b = Bootcamp{
		MyLat:  35.43,
		MyLon:  -1883.32,
		MyDate: time.Now(),
	}
	fmt.Println(b)

	c := Bootcamp{
		43.22, -11.8, time.Now(),
	}

	fmt.Println(c)
	c.getInfo()
}
