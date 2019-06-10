package main

import (
	"fmt"
)

type Vertext struct {
	Lat  float64
	Long float64
}

func main() {
	population := map[string]int{
		"Las Vegas":   845345,
		"Los Angeles": 34523424,
		"New York":    342234,
	}

	// range will return the key and value when used on a map
	for k, v := range population {
		fmt.Printf("%s %d\n", k, v)
	}

	m := map[string]int{"one": 1, "two": 2}
	fmt.Println(m)

	var n map[string]int
	n = make(map[string]int)
	n["one"] = 1

	fmt.Println(n)

	v := map[string]Vertext{
		"Los Angeles": {3.23, 3.23},
	}
	fmt.Println(v)

	//
	v["San Francisco"] = Vertext{4.34, 32.3}
	fmt.Println(v)
}
