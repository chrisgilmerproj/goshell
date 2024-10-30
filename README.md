# GO Shell

This is a golang package used to imitate piping utilities
in the shell

## Example

```sh
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

	err := goshell.NewCommandChain(envVars).
		X([]goshell.Command{
			{"echo", "Hello, World!"}, // Example command
		}).
		Run()

	if err != nil {
		fmt.Println("Error:", err)
	}
}
```
