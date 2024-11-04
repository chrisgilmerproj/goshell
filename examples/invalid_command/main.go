package main

import (
	"log"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	_, err := (&goshell.CommandChain{}).
		Run([]goshell.Command{
			{"invalid_command"},
		})

	if err != nil {
		log.Printf("Expected error running command chain: %v", err)
	}
}
