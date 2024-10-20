package formatter

import (
	"testing"
)

// TestPrintTableHeaders checks the rendering of table headers
func TestPrintTableHeaders(t *testing.T) {

	tbl := InitTable()
	// Render the table to the buffer
	tbl.Render()

	// Print the buffer content to stdout (for visual inspection)

}
