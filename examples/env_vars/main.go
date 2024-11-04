package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	os.Setenv("HELLO_WORLD", "hello, world!")
	defer os.Unsetenv("HELLO_WORLD")

	output, err := (&goshell.CommandChain{}).
		X([]goshell.Command{
			{"bash", "-c", "echo $MY_ENV_VAR"},
		}).
		Run()

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
	fmt.Println(output)
}
