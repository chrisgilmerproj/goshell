package main

import (
	"log"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	err := (&goshell.CommandChain{}).
		X([]goshell.Command{
			{"invalid_command"},
		}).
		Run()

	if err != nil {
		log.Printf("Expected error running command chain: %v", err)
	}
}
