package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("a message")
	log.Fatalln("fatal message")
	log.Panicln("panic message")
}
