package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/ShohamBit/traceectl/pkg/mock"
	"github.com/ShohamBit/traceectl/pkg/models"
)

var rootTests = []models.TestCase{
	{
		Name:           "No root subcommand",
		Args:           []string{"root"},
		ExpectedOutput: rootCmd.Help(), // Update expected output
	},
}

func TestRootCmd(t *testing.T) {
	// Start the mock server
	mockServer, err := mock.StartMockServiceServer()
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer mockServer.Stop() // Ensure the server is stopped after the test

	// Wait for the server to start
	time.Sleep(100 * time.Millisecond)

	for _, test := range rootTests {
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
