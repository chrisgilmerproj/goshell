package main

import (
	"log"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	err := (&goshell.CommandChain{}).
		X([]goshell.Command{
			{"echo", "Hello, World!"},
			{"tr", "[A-Z]", "[a-z]"},
		}).
		Run()

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
}
