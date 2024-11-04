package main

import (
	"fmt"
	"log"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	output, err := (&goshell.CommandChain{}).Run([][]string{
		{"echo", "Hello, World!"},
		{"tr", "[A-Z]", "[a-z]"},
	})

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
	fmt.Println(output)
}
