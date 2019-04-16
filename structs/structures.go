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
	fmt.Println (Bootcamp {
		MyLat: 35.43,
		MyLon: -1883.32,
		MyDate: time.Now(),
	})
}