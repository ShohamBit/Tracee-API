package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/ShohamBit/traceectl/pkg/mock"
	"github.com/ShohamBit/traceectl/pkg/models"
)

var eventTests = []models.TestCase{
	{
		Name:           "No enable events",
		Args:           []string{"event", "enable"},
		ExpectedOutput: "Error: requires at least 1 arg(s), only received 0\n", // Update expected output
	},
	{
		Name:           "Single enable event",
		Args:           []string{"event", "enable", "event1"},
		ExpectedOutput: "Enabled event: event1\n",
	},
	{
		Name:           "Multiple enable events",
		Args:           []string{"event", "enable", "event1", "event2"},
		ExpectedOutput: "Enabled event: event1\nEnabled event: event2\n",
	},

	{
		Name:           "No disable events",
		Args:           []string{"event", "disable"},
		ExpectedOutput: "Error: requires at least 1 arg(s), only received 0\n", // Update expected output
	},
	{
		Name:           "Single disable event",
		Args:           []string{"event", "disable", "event1"},
		ExpectedOutput: "Disabled event: event1\n",
	},
	{
		Name:           "Multiple disable events",
		Args:           []string{"event", "disable", "event1", "event2"},
		ExpectedOutput: "Disabled event: event1\nDisabled event: event2\n",
	},
}

func TestEvent(t *testing.T) {
	// Start the mock server
	mockServer, err := mock.StartMockServer()
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer mockServer.Stop() // Ensure the server is stopped after the test

	// Wait for the server to start
	time.Sleep(100 * time.Millisecond)

	for _, test := range eventTests {
		t.Run(test.Name, func(t *testing.T) {
			// Capture output
			var buf bytes.Buffer
			rootCmd := GetRootCmd()
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			// Set arguments for the test
			rootCmd.SetArgs(test.Args)

			// Execute the command
			if err := rootCmd.Execute(); err != nil {
				t.Error(t, err)
			}

			// Validate output and error (if any)
			output := buf.String()

			if output != test.ExpectedOutput {
				t.Errorf("Expected output: %s, got: %s", test.ExpectedOutput, output)
			}
		})
	}
}
