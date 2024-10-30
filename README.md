# GO Shell

This is a golang package used to imitate piping utilities
in the shell

## Example

```sh
package main

import github.com/chrisgilmerproj/goshell

func main() {
	envVars := map[string]string{
		"MY_ENV_VAR": "my_value",
	}

  // Example 1: Chaining multiple commands
	err := goshell.NewCommandChain(envVars).
		X([]Command{
			{"cat", "file.txt"},           // cat file.txt
			{"cut", "-d", ",", "-f", "1"}, // cut -d"," -f1
			{"grep", "pattern"},           // grep pattern
			{"xargs", "echo"},             // xargs echo
		}).
		Run()

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Example 2: standalone command using a single Command
	err = goshell.NewCommandChain(envVars).
		X([]Command{
			{"mkdir", "newdir"}, // mkdir newdir
		}).
		Run() // Execute the standalone command

	if err != nil {
		fmt.Println("Error:", err)
	}
}
```
