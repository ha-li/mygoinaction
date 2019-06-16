package main

import ("fmt")

type Artist struct {
	Count int
	Album string
}

func immutable (a Artist) int {
	a.Count ++
	a.Album = "A new Album"
	return a.Count
}

func mutable(a *Artist) {
	a.Count ++
	a.Album = "A new Album"
}

func main() {
	a := Artist {0, "My Greatest Hits"}
	fmt.Println(a)
	sold1 := immutable(a)
	fmt.Println(a)
	fmt.Println(sold1)

	mutable(&a)

	// this is now changed
	fmt.Println(a)

	// we can get the pointer from initialization
	b := &Artist {1, "Breaking Ground"}
	mutable(b)
	fmt.Println(b)
}
