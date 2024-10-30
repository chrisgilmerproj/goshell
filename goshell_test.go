package goshell

import (
	"os"
	"testing"
)

func TestCommandChain(t *testing.T) {
	tests := []struct {
		name    string
		command []Command
		wantErr bool
	}{
		{
			name: "SingleCommand",
			command: []Command{
				{"echo", "Hello, World!"},
			},
			wantErr: false,
		},
		{
			name: "MultipleCommands",
			command: []Command{
				{"echo", "Hello, World!"},
				{"tr", "[A-Z]", "[a-z]"},
			},
			wantErr: false,
		},
		{
			name: "EnvironmentVariables",
			command: []Command{
				{"bash", "-c", "echo $MY_ENV_VAR"},
			},
			wantErr: false,
		},
		{
			name: "InvalidCommand",
			command: []Command{
				{"invalid_command"},
			},
			wantErr: true,
		},
		{
			name: "CombinedCommands",
			command: []Command{
				{"echo", "Hello, World!"},
				{"grep", "World"},
			},
			wantErr: false,
		},
		{
			name: "OutputCapture",
			command: []Command{
				{"echo", "Test Output"},
			},
			wantErr: false,
		},
		{
			name: "NestedCommands",
			command: []Command{
				{"echo", "Nested Command"},
				{"tr", "[a-z]", "[A-Z]"},
			},
			wantErr: false,
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
			err := c.X(tt.command).Run()

			// Check for expected error
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}
