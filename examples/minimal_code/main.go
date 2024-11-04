package main

import (
	"fmt"
	"log"

	gs "github.com/chrisgilmerproj/goshell" // Alias the package as 'gs'
)

var CC = &gs.CommandChain{} // CommandChain shortcut

func main() {
	// Using [][]string instead of []gs.Command
	output, err := CC.X([][]string{
		{"echo", "Hello, World!"},
		{"tr", "[A-Z]", "[a-z]"},
	}).Run()

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
	fmt.Println(output)
}
