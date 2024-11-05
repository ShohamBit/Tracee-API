package formatter

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	FormatJSON  = "json"
	FormatTable = "table"
	FormatGoTpl = "gotemplate"
)

// SupportedFormats is a slice of all supported format types
var SupportedFormats = []string{FormatJSON, FormatTable, FormatGoTpl}

type Formatter struct {
	Format string
	Output string
	CMD    *cobra.Command
}

func New(format string, output string, cmd *cobra.Command) (*Formatter, error) {
	if !containsFormat(format) {
		return nil, fmt.Errorf("format %s is not supported", format)

	}
	return &Formatter{
		Format: format,
		Output: output,
		CMD:    cmd,
	}, nil
}

// containsFormat checks if a format is in the SupportedFormats slice
func containsFormat(format string) bool {
	for _, f := range SupportedFormats {
		if f == format {
			return true
		}
	}
	return false
}
