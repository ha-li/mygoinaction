package main

import "fmt"

var (
	truth = false
	name  = "Bob"
)

// constants take on the type of their context
// constants cannot be initialized with :=
const (
	Truth = true
	Name string = "Billy"
	// Religion := "nope"
	Status_OK = 200
	Status_Created = 201
)

func main() {

	// a constant is declared like a variable except with the keyword const
	// instead of var
	var Pi = 3.14
	const pi = 3.14
	// constants can not be initialized with :=
	// const religion := "Still nope"

	fmt.Println(Pi)
    fmt.Println(pi)

	fmt.Println (truth)
	fmt.Println(Truth)

	// print and println are built in, but is is better
	// to use fmt built ins, because of pretty printing
	// for example this will printon +3.140000e+000
	// versus Println will print 3.14
	println(pi)
	print(pi)
	println()

	// use fmt.Sprintf for a String print, ie writes into a string
	var pis = fmt.Sprintf("Pi is %f", pi)
	fmt.Println(pis)

	// use fmt.Printf for formatted printing
	fmt.Printf ("Pi is %f\n", pi)
}
