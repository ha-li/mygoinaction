package main

import "fmt"

func seeall() (int,int,int){
	//1 + 2 + 3
	return 1, 2, 3
}


func main() {
	x, y, z := seeall()

	s := fmt.Sprintf("%d %d %d", x, y, z)
	fmt.Println(s)
}