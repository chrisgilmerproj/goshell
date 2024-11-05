# GO Shell

<img src="./goshell.png" width="128">

A golang package used to imitate piping utilities in the shell.

## Example

```sh
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chrisgilmerproj/goshell"
)

func main() {
	// Set env vars in the environment
	os.Setenv("HELLO_WORLD", "hello, world!")
	defer os.Unsetenv("HELLO_WORLD")

	// Set env vars in the command chain
	CC := goshell.NewCommandChain(map[string]string{"ANOTHER_WORLD": "another, world!"})

	output, err := CC.Run([][]string{
		{"bash", "-c", "echo $HELLO_WORLD"},
	})

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
	fmt.Println(output)

	output, err = CC.Run([][]string{
		{"bash", "-c", "echo $ANOTHER_WORLD"},
	})

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
	fmt.Println(output)
}
```
