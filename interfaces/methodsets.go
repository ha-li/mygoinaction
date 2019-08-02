package main

import (
	"fmt"
)

type Object interface {
	ToString () string
}

type Test interface {
	DoTest () string
}

type Thing struct {
   Name string
   Size int
}

func (t Thing) ToString() string {
   return fmt.Sprintf("Name: %v, Size %d\n", t.Name, t.Size)
}

func (t *Thing) DoTest() string {
	return fmt.Sprintf("Name: %v, Size %d\n", t.Name, t.Size)
}

func runTest(t Test) {
	t.DoTest()
}

func runToString(t Object) {
	t.ToString()
}

func main () {
	t := Thing {"Baseball", 12}
	fmt.Println (t.ToString())
	fmt.Println( (&t).ToString())

	fmt.Println (t.DoTest())
	fmt.Println ((&t).DoTest())

	// this is not okay -- because values of types do not implement
	// an interface when the interface is implemented with a pointer
	// receiver -- due to the method set
	//runTest(t)
	runTest(&t)

	runToString(t)
	runToString(&t)
}