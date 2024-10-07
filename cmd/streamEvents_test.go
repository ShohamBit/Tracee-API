package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/ShohamBit/TraceeClient/mock"
	"github.com/ShohamBit/TraceeClient/models"
)

var StreamEventTests = []models.TestCase{
	{
		Name: "No Policies",
		Args: []string{"streamEvents"}, // Removed the empty string argument
		ExpectedOutput: `policies:{matched:""}
policies:{matched:"policy1"}
policies:{matched:"policy2"}
policies:{matched:"policy3"}
policies:{matched:"policy1" matched:"policy3"}
policies:{matched:"policy1" matched:"policy2"}
policies:{matched:"policy2" matched:"policy3"}
policies:{matched:"policy1" matched:"policy2" matched:"policy3"}
`,
	},
	{
		Name: "Single policy",
		Args: []string{"streamEvents", "policy1"},
		ExpectedOutput: `policies:{matched:"policy1"}
policies:{matched:"policy1" matched:"policy3"}
policies:{matched:"policy1" matched:"policy2"}
policies:{matched:"policy1" matched:"policy2" matched:"policy3"}
`,
	},
	{
		Name: "Multiple policies",
		Args: []string{"streamEvents", "policy1", "policy2"},
		ExpectedOutput: `policies:{matched:"policy1"}
policies:{matched:"policy2"}
policies:{matched:"policy1" matched:"policy3"}
policies:{matched:"policy1" matched:"policy2"}
policies:{matched:"policy2" matched:"policy3"}
policies:{matched:"policy1" matched:"policy2" matched:"policy3"}
`,
	},
}

func TestStreamEvent(t *testing.T) {
	for _, test := range StreamEventTests { // Corrected the variable name
		t.Run(test.Name, func(t *testing.T) {
			// Start the mock server
			mockServer, err := mock.StartMockServiceServer()
			if err != nil {
				t.Fatalf("Failed to start mock server: %v", err)
			}
			defer mockServer.Stop() // Ensure the server is stopped after the test

			// Wait for the server to start
			time.Sleep(100 * time.Millisecond)

			// Capture output
			var buf bytes.Buffer
			rootCmd.SetOut(&buf)
			rootCmd.SetErr(&buf)

			// Set arguments for the test
			rootCmd.SetArgs(test.Args)

			// Execute the command
			if err := rootCmd.Execute(); err != nil {
				t.Fatalf("Execute() failed: %v", err)
			}

			// Check if output matches expected output
			if buf.String() != test.ExpectedOutput {
				t.Errorf("Expected %q, \ngot %q", test.ExpectedOutput, buf.String())
			}
		})
	}
}
