package main

import (
	"fmt"
	"log"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	// Use your library
	envVars := map[string]string{
		"MY_ENV_VAR": "my_value",
	}

	output, err := goshell.NewCommandChain(envVars).
		Run([][]string{
			{"echo", "Hello, World!"}, // Example command
		})

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}

	fmt.Println(output)
}
