package main

import (
	"fmt"
	s "strings"
	//"log"
	//"os"
)

type Block struct {
	AttachmentPoint float64
	JointCounts int
}



// a method that looks like a constructor
// but it does not belong to Block
// ie it is not
//      func (b Block) NewBlock ...
func NewBlock (p float64, jc int) *Block {
	// new(T) is a builtin that returns a pointer to T
	// where T is a type
	b:=new(Block)
	b.AttachmentPoint = p
	b.JointCounts = jc
	return b
}

// a method that functions like a constructor
func MakeBlock(p float64) Block {
	// make only works on channels, slices and array
	b := Block {p, 4}
	return b
}

func (b Block) DoesNotRotate180() {
	//fmt.Println(b.AttachmentPoint)
	b.AttachmentPoint = b.AttachmentPoint - 180
	//fmt.Println(b.AttachmentPoint)
}

func (b *Block) Rotate180() {
	//fmt.Println(b.AttachmentPoint)
	b.AttachmentPoint = b.AttachmentPoint - 180
	//fmt.Println(b.AttachmentPoint)
}

func (b Block) DoNotMutateJointCount(jc int) {
	b.JointCounts = jc
}

func (b *Block) MutateJointCount(jc int) {
	b.JointCounts = jc
}

// this is the proper way to build composition in Golang
type Blocker struct {
	Block
	FunnyString string
}

// this is NOT how you want to build composition in GoLang
// see if DudBlocker inherits the methods of Block
type DudBlocker struct {
	Dud Block   // this type has a name, so it is not composition
	FunnyString string
}

type Socker struct {
	FunniestString string

	// the order of the types does not matter
	Block
}

type FakeString struct {
	string   // this is an invalid composition, will not compile
}

func main() {
	b := NewBlock(4.56, 12)
	// this will print out & to reflect its a pointer
	fmt.Println(b)

	// this will not print &
	c := MakeBlock(6.5)
	fmt.Println(c)

	c.Rotate180()
	fmt.Println(c)

	c.DoesNotRotate180()
	fmt.Println(c)

	c.MutateJointCount(22)
	fmt.Println(c)

	c.DoNotMutateJointCount(22)
	fmt.Println(c)

	b2 := Blocker{
		Block{AttachmentPoint: 5.43, JointCounts: 11},
		"Blocker",
	}

	// here the methods of block are also methods of blocker
	b2.MutateJointCount(1)

	d2 := DudBlocker{
		Block{AttachmentPoint: 2.1, JointCounts: 18},
		"DudBlocker",
	}
	fmt.Println(d2)

	// here the methods of block are NOT methods of dudblocker
	// this does not compile
	// because of the composition not being correct
	d2.MutateJointCount(1)

	s1 := Socker{
		"Booyah",
		Block{AttachmentPoint: 5.43, JointCounts: 11},
	}

	s1.MutateJointCount(44)
	fmt.Println(s1)

	s3 := "RealString"
	fmt.Println(s.ToUpper(s3))

	s2 := FakeString {"FakeString"}
	// fmt.Println(s2)
	fmt.Println(s.ToUpper(s2))
}