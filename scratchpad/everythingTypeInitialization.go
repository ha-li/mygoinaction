package main

import (
	"fmt"
)

// the different ways to initialize a type in go
// 1. dot notation
// 2. struct literal
// 3. Name: syntax (allows you to partially initialize a type
// 4. a 'constructor' function
// 5. new and dot notation

type MysteryType struct {
	Mutant bool
	Name string
}

func NewMysteryType(isMutant bool, name string) *MysteryType {
	return &MysteryType {
		Mutant: isMutant,
		Name: name,
	}
}

func main () {

	// struct literal using Name syntax
	m1 := MysteryType {
   	   Mutant: true,
   	   Name: "apple",
    }

   fmt.Println(m1)

	// two examples of Name syntax with only a partial list initialization
	m1b := MysteryType {
		Name: "jeepers",
	}
	fmt.Println(m1b)

   m2 := MysteryType {
   	false, "jessibel",
   }

   fmt.Println(m2)

   // dot notation
   m3 := MysteryType{}
   m3.Mutant = false
   m3.Name = "nathan"

   fmt.Println(m3)

   // a 'constructor' function
   m4 := NewMysteryType(false, "Jessica")
   fmt.Println(m4)

   // new and dot notation
   m5 := new(MysteryType)
   m5.Mutant = false
   m5.Name = "daemon"
   fmt.Println(m5)
}