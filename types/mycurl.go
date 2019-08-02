package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


func init () {
	if len(os.Args) != 2 {
		fmt.Printf( "Usage: ./mycurl <url>\n")
        os.Exit(-1)
	}
}

func main () {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy (os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Booyah")
}
