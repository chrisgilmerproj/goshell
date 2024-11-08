# GO Shell

<img src="./goshell.png" width="128">

A golang package used to imitate piping utilities in the shell.

## Motivation

The motivation here is to have a way to convert bash scripts
to golang while not having to implement things immediately in
native golang or using SDKs. For example, moving a complex
bash script of several hundred lines of code might be difficult
if you try to replace all the bash utilities. This is especially
difficult if you have chained commands which pipe from one
command to another. The library fills the need for having an
easy-to-use replacement for piping commands that is portable
across multiple golang scripts.

## Example

If you were using bash you might have a script that includes
setting env vars and piping. This contrived example gives a
good example of what you might encounter:

```sh
#! /usr/bin/env bash

set -euo pipefail

HELLO_WORLD="hello, world!" bash -c 'echo $HELLO_WORLD' | tr '[:lower:]' '[:upper:]'
```

The output would be: `HELLO, WORLD!`.

The equivalent in golang would look like:

```sh
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
		{"tr", "[:lower:]", "[:upper:]"},
	})

	if err != nil {
		log.Fatalf("Error running command chain: %v", err)
	}
	fmt.Println(output)
}
```

For obvious reasons you wouldn't convert a tiny script like
this example to golang. However, if you had a long bash script
and converting it required maintaining chained commands then
you might find this to be a helpful utility. Especially if the
tools you are working with are easier to shell out to instead of
using native golang SDKs.
