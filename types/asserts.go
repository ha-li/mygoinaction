package main

import (
	"fmt"
	"time"
)

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

// showing a simple type assertion
func test() {
	foo := map[string]interface{}{
		"Matt": 42,
	}

	timeMap(foo)
	fmt.Println(foo["updated_at"])
	fmt.Println(foo)
}

func doAsserts() {
   var i interface{} = "string"
   fmt.Println(i)
}

func doStringTypeAssert(y interface{}) {
	z, ok := y.(string)
	if ok {
		fmt.Println(len(z))
	} else {
		fmt.Printf("%v is not a string\n", y)
	}
}

func showDynamic(x interface{}) {
	// now x is assigned an int, ok due to declaration
	x = 7
	fmt.Println(x)

	// x is assigned a string, still ok
	x = "random string"
	fmt.Println(x)
}


func main() {
	test()
	doAsserts()
	doStringTypeAssert("a simple string")
	doStringTypeAssert(3)

	// a starts off as a float, but the declaration says
	// it's type may change
	var a interface{} = 4.34
	showDynamic(a)

	// b is declared as a float64, but it still may change in showDynamic
	var b float64 = 2.1
	showDynamic(b)
}