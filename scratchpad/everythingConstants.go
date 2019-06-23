package main

import (
	"fmt"
)

// constants get their own block
const (
	Pi = 3.14
	OK = 200
	HTTP_OK = "OK"
)

func main() {
	fmt.Println(OK)
	fmt.Println(Pi*3*3)
	fmt.Println (HTTP_OK)
}