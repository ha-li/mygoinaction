package main

import (
	"fmt"
	"net/url"
	"time"
)

func main () {
	fmt.Println("hello")
	google,_ := url.QueryUnescape("www.google.com")
	fmt.Println(google)
	t1 := time.Now()
	fmt.Println("t1: ", t1)
	t2 := t1.Add(40 * time.Second)
	fmt.Println("t1: ", t1)
	fmt.Println("t2: ", t2)
}