package main

import (
	"fmt"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	// Use your library
	envVars := map[string]string{
		"MY_ENV_VAR": "my_value",
	}

	output, err := goshell.NewCommandChain(envVars).
		Run([]goshell.Command{
			{"echo", "Hello, World!"}, // Example command
		})

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(output)
}
