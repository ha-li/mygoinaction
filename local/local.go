package local

// custom packages should be named
// as the last element of the import path
// so this path is ...local, so the
// package is declared local

import (
	"fmt"
)

func Locate (n string) {
   fmt.Println(n)
}
