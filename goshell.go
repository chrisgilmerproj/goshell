package goshell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CommandChain struct to store output, errors, and environment variables
type CommandChain struct {
	output       []byte
	err          error
	envVariables []string
	commands     [][]string
}

// NewCommandChain creates a new CommandChain with optional environment variables
func NewCommandChain(envVars map[string]string) *CommandChain {
	// Start with current environment variables
	env := os.Environ()

	// Create a map from the current environment variables
	envMap := make(map[string]string)
	for _, e := range env {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			envMap[pair[0]] = pair[1]
		}
	}

	// Override with the explicit environment variables
	for key, value := range envVars {
		envMap[key] = value
	}

	// Convert the merged environment variable map back to a slice
	var envSlice []string
	for key, value := range envMap {
		envSlice = append(envSlice, fmt.Sprintf("%s=%s", key, value))
	}

	return &CommandChain{
		envVariables: envSlice,
	}
}

// Run executes the commands in the chain
func (cc *CommandChain) Run(commands [][]string) (string, error) {
	var input *bytes.Buffer // Start with no input
	cc.commands = commands  // Set the commands

	for _, cmd := range cc.commands {
		// Create command and set up input if necessary
		command := exec.Command(cmd[0], cmd[1:]...) // Create command
		command.Env = cc.envVariables               // Set environment variables

		// If there is input from the previous command, set it as the command's stdin
		if input != nil {
			command.Stdin = input
		}

		// Capture the output
		var stdoutBuf bytes.Buffer
		var stderrBuf bytes.Buffer
		command.Stdout = &stdoutBuf
		command.Stderr = &stderrBuf

		// Execute the command
		if err := command.Run(); err != nil {
			return stderrBuf.String(), err // Return error if command fails
		}

		// The output of the current command becomes the input for the next
		input = &stdoutBuf
	}

	// Return the output from the last command
	return input.String(), nil
}
