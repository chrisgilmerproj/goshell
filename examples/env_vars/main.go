package main

import (
	"log"
	"os"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	os.Setenv("MY_ENV_VAR", "my_value")
	defer os.Unsetenv("MY_ENV_VAR")

	err := (&goshell.CommandChain{}).
		X([]goshell.Command{
			{"bash", "-c", "echo $MY_ENV_VAR"},
		}).
		Run()

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
}
