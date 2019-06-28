package main

import (
	"fmt"
)

// a simple practice of the for loop iterating through each element
// ['a'] has an index of 0 and length of 1, starting at 0... the last element is at length - 1
// which means if you are using < in your loop construct, it is < len(s)
// and if you are using <= in the loop construct, it is <= len(s) - 1

// ['a', 'b'] the index is 0, 1, and the length in 2, so from 0...1, the last index to access is len - 1
// so your loop construct is i:= 0; i < len(s); i++
// or i:=0; i <= len(s) - 1; i ++
func main() {
	s := [1]byte { 'a' }
	for i:=0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
	fmt.Println()

	s2 := [2]byte { 'a', 'b' }
	for i:=0; i < len(s2); i++ {
		fmt.Println(string(s2[i]))
	}
}
