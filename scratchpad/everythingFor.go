package main

import (
	"fmt"
)

func for1 () {

   i := 0

   // this is an infinite loop
   for {
   	   // if statements have a short statement before
   	   // the if clause, here it is i++;
   	   // the if clause is the test i == 1000
   	   if i++; i == 1000 {
   	   	  break
	   }

   	   fmt.Println(i)
   }
}

func for2 () {
	i := 0

	// this is a simple for loop with just a test
	// this is the same as a while loop
	for i < 10 {
		i++
		fmt.Println (i)
	}
}

func for3 () {
	i := 0

	// a basic for loop with no pre and post statement
	for ; i < 10; {
		i++
		fmt.Println (i)
	}

	// a tradition for loop with pre and post statements
	for j:=0; j < 10; j++ {
		fmt.Println(j)
	}
}


func main() {
	for1()
}
