package goshell

import (
	"os"
	"testing"
)

func TestCommandChain(t *testing.T) {
	tests := []struct {
		name    string
		command [][]string
		wantErr bool
		wantOut string
	}{
		{
			name: "SingleCommand",
			command: [][]string{
				{"echo", "Hello, World!"},
			},
			wantErr: false,
			wantOut: "Hello, World!\n",
		},
		{
			name: "MultipleCommands",
			command: [][]string{
				{"echo", "Hello, World!"},
				{"tr", "[A-Z]", "[a-z]"},
			},
			wantErr: false,
			wantOut: "hello, world!\n",
		},
		{
			name: "EnvironmentVariables",
			command: [][]string{
				{"bash", "-c", "echo $MY_ENV_VAR"},
			},
			wantErr: false,
			wantOut: "override_value\n",
		},
		{
			name: "InvalidCommand",
			command: [][]string{
				{"invalid_command"},
			},
			wantErr: true,
			wantOut: "",
		},
		{
			name: "CombinedCommands",
			command: [][]string{
				{"echo", "Hello, World!"},
				{"grep", "World"},
			},
			wantErr: false,
			wantOut: "Hello, World!\n",
		},
		{
			name: "OutputCapture",
			command: [][]string{
				{"echo", "Test Output"},
			},
			wantErr: false,
			wantOut: "Test Output\n",
		},
		{
			name: "NestedCommands",
			command: [][]string{
				{"echo", "Nested Command"},
				{"tr", "[a-z]", "[A-Z]"},
			},
			wantErr: false,
			wantOut: "NESTED COMMAND\n",
		},
	}

	// Set environment variable for tests
	os.Setenv("MY_ENV_VAR", "my_value")
	defer os.Unsetenv("MY_ENV_VAR") // Clean up after the test

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCommandChain(nil)

			// Special handling for environment variable test case
			if tt.name == "EnvironmentVariables" {
				// Override with an environment variable
				c = NewCommandChain(map[string]string{"MY_ENV_VAR": "override_value"})
			}

			// Execute the command chain
			out, err := c.Run(tt.command)

			// Check for expected error
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error in test %s: %v, got: %v, stderr: %s", tt.name, tt.wantErr, err, out)
			}
			if out != tt.wantOut {
				t.Errorf("expected output in test %s: %s, got: %s", tt.name, tt.wantOut, out)
			}
		})
	}
}
