package main

import (
	"fmt"
	"log"

	"github.com/chrisgilmerproj/goshell"
)

func main() {

	// Set env vars in the command chain
	CC := goshell.NewCommandChain(map[string]string{
		"HELLO_WORLD": "hello, world!",
	})

	output, err := CC.Run([][]string{
		{"bash", "-c", "echo $HELLO_WORLD"},
		{"tr", "a-z", "A-Z"},
	})

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
	fmt.Println(output)
}
