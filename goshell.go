package goshell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Command type as a slice of strings
type Command []string

// CommandChain struct to store output, errors, and environment variables
type CommandChain struct {
	output       []byte
	err          error
	envVariables []string
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

// X function to execute a list of commands. X is for eXecute.
func (c *CommandChain) X(commands []Command) *CommandChain {
	// If there's already an error, skip execution
	if c.err != nil {
		return c
	}

	// Execute each command in the list
	for _, cmd := range commands {
		// Prepare the command
		execCmd := exec.Command(cmd[0], cmd[1:]...)
		execCmd.Env = c.envVariables // Set the environment variables

		// Pass the output from the previous command as stdin to the next
		if c.output != nil {
			execCmd.Stdin = bytes.NewReader(c.output)
		}

		// Run the command and capture the output and errors
		c.output, c.err = execCmd.Output()
		if c.err != nil {
			return c // Exit if there's an error
		}
	}

	return c
}

// Run final output or return error
func (c *CommandChain) Run() error {
	// Return error if any occurred during the chain
	if c.err != nil {
		return c.err
	}

	// Print the final output
	fmt.Println(string(c.output))
	return nil
}
