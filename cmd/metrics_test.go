package cmd

import (
	"bytes"
	"testing"
	"time"

	"github.com/ShohamBit/TraceeClient/mock"
	"github.com/stretchr/testify/assert"
)

func TestMetricsCommand(t *testing.T) {
	// Start the mock server
	mockServer, err := mock.StartMockDiagnosticServer()
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	defer mockServer.Stop() // Ensure the server is stopped after the test

	// Wait for a moment to ensure the server is up
	time.Sleep(100 * time.Millisecond)

	// Capture output
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)

	// Set the arguments for the metrics command
	rootCmd.SetArgs([]string{"metrics"})

	// Execute the metrics command
	err = rootCmd.Execute()
	assert.NoError(t, err, "Expected no error when executing metrics command")

	// Expected output (adjust as per mock server's response)
	expectedMetrics := []string{
		"EventCount: 1",
		"EventsFiltered: 2",
		"NetCapCount: 3",
		"BPFLogsCount: 4",
		"ErrorCount: 5",
		"LostEvCount: 6",
		"LostWrCount: 7",
		"LostNtCapCount: 8",
		"LostBPFLogsCount: 9",
	}

	// Get the output from the buffer
	output := buf.String()

	// Assert that all expected metrics are in the output
	for _, metric := range expectedMetrics {
		assert.Contains(t, output, metric, "Output should contain metric: "+metric)
	}
}
