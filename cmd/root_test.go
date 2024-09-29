package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {

	// Capture output
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	// Execute the root command
	rootCmd.SetArgs([]string{})
	rootCmd.Execute()

	// Check if help was printed
	expectedHelpPrefix := "Tracee client is the client for tracee api server."
	output := buf.String()
	assert.Contains(t, output, expectedHelpPrefix, "Root command help output should contain the expected prefix")

}
