package formatter

import (
	"fmt"
	"strings"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

type Formatter struct {
	format string
	output string
	cmd    *cobra.Command
}

func New(format string, output string, cmd *cobra.Command) *Formatter {
	return &Formatter{
		format: format,
		output: output,
		cmd:    cmd,
	}
}
func (f *Formatter) PrintTableHeaders() {
	f.cmd.Printf("%-15s %-20s %-16s %-25s %s\n",
		"TIME",
		"NAME",
		"POLICIES",
		"CONTEXT",
		"DATA",
	)
}
func (f *Formatter) PrintTableRow(event *pb.Event) {
	timestamp := event.Timestamp.AsTime().Format("15:04:05.000")

	f.cmd.Printf("%-15s %-20s %-15s %-25s %s\n",
		timestamp,
		event.Name,
		event.Policies.Matched,
		event.Context,
		getEventData(event.Data),
	)

}

// func getEventContext(context *pb.Context) string {
// 	return " "
// }

// generate event data
func getEventData(data []*pb.EventValue) string {
	var result []string
	for _, ev := range data {
		result = append(result, getEventName(ev)+getEventValue(ev))
	}
	return strings.Join(result, ", ")
}
func getEventName(ev *pb.EventValue) string {
	return strings.ToUpper(ev.Name[0:1]) + ev.Name[1:] + ": "
}

// generate event value
func getEventValue(ev *pb.EventValue) string {
	switch v := ev.Value.(type) {
	case *pb.EventValue_Int32:
		return fmt.Sprintf("%d", v.Int32)
	case *pb.EventValue_Int64:
		return fmt.Sprintf("%d", v.Int64)
	case *pb.EventValue_UInt32:
		return fmt.Sprintf("%d", v.UInt32)
	case *pb.EventValue_UInt64:
		return fmt.Sprintf("%d", v.UInt64)
	case *pb.EventValue_Str:
		return v.Str
	case *pb.EventValue_Bytes:
		return fmt.Sprintf("%x", v.Bytes)
	case *pb.EventValue_Bool:
		return fmt.Sprintf("%t", v.Bool)
	case *pb.EventValue_StrArray:
		return strings.Join(v.StrArray.Value, ", ")
	case *pb.EventValue_Int32Array:
		return fmt.Sprintf("%v", v.Int32Array.Value)
	case *pb.EventValue_UInt64Array:
		return fmt.Sprintf("%v", v.UInt64Array.Value)
		//TODO: add more types
	default:
		// if data type not supported yet
		return "unknown"
	}
}
