package main

import (
	"fmt"
	s "strings"
	"bytes"
)


// a function called p1, that returns a function that returns an in
func p1() (func() int) {
	// a free variable
	i := 5

	// an anonymous function that returns an int
	return func() int {

		// prints the free variable then increments it and returns i
		fmt.Println (i)
		i++
		return i
    }
}

// another syntax of a function that returns a function
// that takes a string param & returns a string
// note the syntax of the return value
func p2() func(c string) string {
	s1 := "hello"

	return func (c string) string {
		// prints out if hello contains c then returns hello
		//fmt.Println(s1)
		fmt.Println(s.Contains(s1, c))

		// append c to s1, but this will not change s1 in sub calls
		var buffer bytes.Buffer
		buffer.WriteString(s1)
		buffer.WriteString(c)
		s1 := buffer.String()
		//fmt.Println(s1)
		return s1
	}
}


func main () {

	// get an instance of the anonymous function
	// where i = 5, assign it to f
	// and call it 3 times, each time incrementing i
	f := p1()
	j := f()
	k := f()
	l := f()
	fmt.Printf("%d %d %d\n", j, k, l)

	// get another instance of the anonymous function,
	// where i = 5, assign it to g, and call g once
	g := p1()
	m := g()
	fmt.Printf ("%d\n", m)

	h := p2()
	h("c")
	h("c")

	h("h")

}