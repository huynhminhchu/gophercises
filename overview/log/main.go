package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: hello world")
	} else {
		log.Panic("Patic: goodbye world")
	}
}
