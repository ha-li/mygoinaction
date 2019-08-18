package main

import ("fmt")

func send2 (c chan int, v int) {
	c <- v
}

func main () {
	c := make (chan int)
	go send2 (c, 5)

	fmt.Println(<-c)
}
