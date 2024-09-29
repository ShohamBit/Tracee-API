package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/ShohamBit/TraceeClient/mock"
)

/* data source client */
/* diagnostic client */
/* service client */
var (
	ExpectedVersion = "Version: v0.22.0-15-gd09d7fca0d\n" // Match the output format

)

func TestVersionCommand(t *testing.T) {
	// Start the mock server
	mockServer, err := mock.StartMockServiceServer()
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer mockServer.Stop() // Ensure the server is stopped after the test

	// Wait for a moment to ensure the server is up
	time.Sleep(100 * time.Millisecond)

	// Capture output
	var buf bytes.Buffer
	//rootCmd := rootCmd
	rootCmd.SetOut(&buf)

	// Execute the version command
	rootCmd.SetArgs([]string{"version"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Check expected output
	if buf.String() != ExpectedVersion {
		t.Errorf("Expected %q, got %q", ExpectedVersion, buf.String())
	}

}
