package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// a buffer variable, better name would be buffer
	var b bytes.Buffer
	// write into the buffer
	b.Write([]byte("Hello"))
	fmt.Fprintf(&b, " World!\n")
	io.Copy(os.Stdout, &b)
}
