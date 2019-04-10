package main

import (
	"fmt"
)

type Artist struct {
	Name string
	Genre string
	Songs int
}

func newRelease (artist Artist) int {
	artist.Songs++
	return artist.Songs
}

func main () {
	elvis := Artist {Name: "Elvis", Genre: "Hipster", Songs: 50}
   fmt.Println (elvis)
	fmt.Printf ("%s released %d songs\n", elvis.Name, newRelease(elvis))
	fmt.Printf ("%s released %d songs\n", elvis.Name, elvis.Songs)
}
