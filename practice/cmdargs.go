package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/goinaction/code/chapter3/words"
)

func main() {
	program := os.Args[0]

	if len(os.Args) == 1 {
		fmt.Printf("Usuage: \n\t%s filename\n", program)
		return
	}

	filename := os.Args[1]
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if true {
		fmt.Println("true")
	}

	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("there are %d words in your text\n", count, 3.2)
}
