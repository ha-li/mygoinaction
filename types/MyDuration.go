package main

import (
	"fmt"
	"time"
)

type MyDuration int64

func (d MyDuration) Change (i int, name string) MyDuration {
	fmt.Printf("address inside Change of %s %v\n", name, &d)
	fmt.Printf ("value of d %d, of i %d\n", d, i)
	itmp := int64(d) + int64(i)
	d = MyDuration(itmp)
	fmt.Printf ("value of d after change %d \n", d)
	return d
}

func main() {
	d := 4

	md := MyDuration(5)

	// the address of md in main
	fmt.Printf ("address of md %v %d\n", &md, md)
	// here md is used as value, calling Change,
	// and in Change, we print out the address of
	// md (ie d) and see that the address is different
	// so md != d, it is copy by value
	md2 := md.Change(d, "md")

	// now md2 is assigned the value of the return value of Change
	// which was d, but that was copy by value, so md2 != d also
	fmt.Printf("address of md %v %d\n", &md, md)

	// confirm that by printing out md2 also
	fmt.Printf("address of md2 %v %d \n", &md2, md2)
}
