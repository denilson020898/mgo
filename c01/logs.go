package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello Fucker!")
	}
	log.Panic("Panic: Hello Panic!")
}
