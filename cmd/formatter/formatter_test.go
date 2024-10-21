package formatter

import (
	"testing"
)

// TestPrintTableHeaders checks the rendering of table headers
func TestPrintTableHeaders(t *testing.T) {

	tbl := InitTable("stdout")
	// Render the table to the buffer
	// Print the buffer content to stdout (for visual inspection)
	tbl.Render()

	/*
		// Expected output
		expected := InitTable("stdout")
		expected.Render()

		var buf bytes.Buffer
		rootCmd.SetOut(&buf)
		rootCmd.SetErr(&buf)
		output := buf.String()
	*/
}
