package cmd_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/ShohamBit/TraceeClient/cmd"
	"github.com/ShohamBit/TraceeClient/pkg/mock"
	"github.com/ShohamBit/TraceeClient/pkg/models"
	"github.com/stretchr/testify/assert"
)

var (
	EnableEventTests = []models.TestCase{
		{
			Name:           "No events",
			Args:           []string{"enableEvent"},
			ExpectedOutput: "Error: requires at least 1 arg(s), only received 0\n", // Update expected output
		},
		{
			Name:           "Single event",
			Args:           []string{"enableEvent", "event1"},
			ExpectedOutput: "Enabled event: event1\n",
		},
		{
			Name:           "Multiple events",
			Args:           []string{"enableEvent", "event1", "event2"},
			ExpectedOutput: "Enabled event: event1\nEnabled event: event2\n",
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

	for _, test := range EnableEventTests {
		t.Run(test.Name, func(t *testing.T) {
			// Capture output
			var buf bytes.Buffer
			rootCmd := cmd.GetRootCmd()
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			// Set arguments for the test
			rootCmd.SetArgs(test.Args)

			// Execute the command
			err := rootCmd.Execute()

			// Validate output and error (if any)
			output := buf.String()

			// If no arguments provided, we expect an error
			if test.Name == "No events" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Check if output matches expected output
			assert.Contains(t, output, test.ExpectedOutput)
		})
	}
}
