package main

import (
	"fmt"
)

func setChan (c chan string) {
	c <- "Hello World"
}


func getChan () chan string {
	return make(chan string)
}

func main () {
	c := getChan()

	go setChan(c)

	fmt.Println (<- c)
}