package main

import (
	"fmt"
	"log"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	output, err := (&goshell.CommandChain{}).
		Run([][]string{
			{"invalid_command"},
		})

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
	fmt.Println(output)
}
