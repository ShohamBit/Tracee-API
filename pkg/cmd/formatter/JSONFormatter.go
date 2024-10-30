package formatter

import (
	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

// PrintJSON prints an event in JSON format
func PrintJSON(cmd *cobra.Command, event *pb.Event, _ string) {
	//TODO: add more output formats
	cmd.Printf("%s\n", event.String())
}
