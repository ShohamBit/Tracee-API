package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/ShohamBit/TraceeClient/mock"
	"github.com/ShohamBit/TraceeClient/models"
	"github.com/stretchr/testify/assert"
)

var (
	DisableEventTests = []models.TestCase{
		{
			Name:           "No events",
			Args:           []string{"disableEvent"},
			ExpectedOutput: "Error: requires at least 1 arg(s), only received 0\n", // Update expected output
		},
		{
			Name:           "Single event",
			Args:           []string{"disableEvent", "event1"},
			ExpectedOutput: "Disabled event: event1\n",
		},
		{
			Name:           "Multiple events",
			Args:           []string{"disableEvent", "event1", "event2"},
			ExpectedOutput: "Disabled event: event1\nDisabled event: event2\n",
		},
	}
)

func TestDisableEvent(t *testing.T) {
	// Start the mock server
	mockServer, err := mock.StartMockServiceServer()
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer mockServer.Stop() // Ensure the server is stopped after the test

	// Wait for the server to start
	time.Sleep(100 * time.Millisecond)

	for _, test := range DisableEventTests {
		t.Run(test.Name, func(t *testing.T) {
			// Capture output
			var buf bytes.Buffer
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
