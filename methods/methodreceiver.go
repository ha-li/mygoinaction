package main

import (
	"fmt"
)

type TestObject struct {
	Key int
	Value string
}

func (t TestObject) AMethod1() {
	fmt.Println (t.Key)
}

func (t *TestObject) AMethod2() {
	// note this is okay
	fmt.Println (t.Key)

	// this is also okay
	fmt.Println ((*t).Key)
}

// this method will mutate t
// but since t is passed by value
// the receiver object will not actually be affected
// see test3
func (t TestObject) AMutator() {
	t.Key = 23
}

func (t *TestObject) AMutator2() {
	t.Key = 99
}

func (t TestObject) AMutator3() int {
	return 100
}

func (t TestObject) AMutator4() TestObject {
	t.Key = 101
	return t
}

func main () {
	fmt.Println( "Hello World")

	// you can use a value to call a method with a pointer receiver or value receiver
	// go compiler will adjust the value to the correct receiver
	// this is a convencience factor go designers did on purpose
	test1 := TestObject {123, "ABCD" }
	test1.AMethod1()
	test1.AMethod2()

	// you can use a pointer to call a method with a pointer receiver or a value reciever
	// go compiler will adjust the value to the correct receiver
	// this is a convencience factor go designers did on purpose
	test2 := &TestObject { 987, "ZYX" }
	test2.AMethod1()
	test2.AMethod2()

	// test3 is created with Key as 1, then mutated by AMutator method to 23
	test3 := TestObject { 1, "A" }
	test3.AMutator()
	// but when we print Key, it is still 1
	test3.AMethod1()

	// the reason test3.Key is still 1 after mutating is because of
	// pass by value, test3 is copied and passed into AMutator, which mutates Key
	// but in main, the original test3 object is not changed,

	// test4 is a value object with Key as 2, it is mutated by AMutator2 method, Key
	// is set to 99
	test4 := TestObject { 2, "B" }
	test4.AMutator2()
	// and then when we print out Key, it is 99
	test4.AMethod1()

	// the reason test4 has mutated is because the Mutator2 method takes a pointer
	// so the object address is copied and passed into Mutator2, which is then
	// able to reference the original object, mutate it, and then when it
	// ends, in main, it is also changed

	// this is another way of mutating a object is to leave it to the
	// caller to save the new value in the receiver if they want to save it
	test5 := TestObject {3, "C" }
	test5.Key = test5.AMutator3()
	test5.AMethod1()

	// this is another example of implementing a mutator by pass by value
	//
	test6 := TestObject { 4, "D" }
	test6.AMutator4()
	test6.AMethod1()

	test6 = test6.AMutator4()
	test6.AMethod1()
}
