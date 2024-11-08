#! /usr/bin/env bash

set -euo pipefail

HELLO_WORLD="hello, world!" bash -c 'echo $HELLO_WORLD' | tr '[:lower:]' '[:upper:]'
