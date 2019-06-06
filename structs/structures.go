package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	MyLat float64
	MyLon float64
	MyDate time.Time
}


func main() {
	var b = Bootcamp {
		MyLat: 35.43,
		MyLon: -1883.32,
		MyDate: time.Now(),
	}
	fmt.Println(b)

	c := Bootcamp {
		43.22, -11.8, time.Now(),
	}

	fmt.Println(c)
}