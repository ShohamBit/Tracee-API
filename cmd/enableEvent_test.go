package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/ShohamBit/TraceeClient/mock"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name           string
	args           []string
	expectedOutput string
}

var (
	tests = []TestCase{
		{
			name:           "Single event",
			args:           []string{"enableEvent", "event1"},
			expectedOutput: "Enabled event: event1\n",
		},
		{
			name:           "Multiple events",
			args:           []string{"enableEvent", "event1", "event2"},
			expectedOutput: "Enabled event: event1\nEnabled event: event2\n",
		},
		{
			name:           "No events",
			args:           []string{"enableEvent"},
			expectedOutput: "Error: requires at least 1 arg(s), only received 0\n", // Update expected output
		},
	}
)

func TestEnableEvent(t *testing.T) {
	// Start the mock server
	mockServer, err := mock.StartMockServiceServer()
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer mockServer.Stop() // Ensure the server is stopped after the test

	// Wait for the server to start
	time.Sleep(100 * time.Millisecond)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Capture output
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			// Set arguments for the test
			rootCmd.SetArgs(test.args)

			// Execute the command
			err := rootCmd.Execute()

			// Validate output and error (if any)
			output := buf.String()

			// If no arguments provided, we expect an error
			if test.name == "No events" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Check if output matches expected output
			assert.Contains(t, output, test.expectedOutput)
		})
	}
}
