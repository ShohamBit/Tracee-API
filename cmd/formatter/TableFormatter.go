package formatter

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	table "github.com/aquasecurity/table"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
)

// init new table
func InitTable(_ string) *table.Table {
	// TODO: support other output formats
	tbl := table.New(os.Stdout)
	tbl.SetLineStyle(table.StyleBold)
	initTableHeaders(tbl)
	return tbl
}

// Define table titles once, for example during table initialization
func initTableHeaders(tbl *table.Table) {
	tbl.SetHeaders(
		"Time",
		"Command",
		"Policy",
		"Context",
		"Data")

	tbl.SetAutoMergeHeaders(true)
}
func PrintTableRow(tbl *table.Table, event *pb.Event) {
	clearTerminal()
	tbl.AddRow(
		event.GetTimestamp().AsTime().Format("15:04:05.000"), // time
		event.Name, // command
		strings.Join(event.Policies.Matched, ", "), // policy
		event.Context.String(),                     // context
		getEventData(event.Data),
	)
	tbl.Render()
}
func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout // Set output to standard output
	cmd.Run()
}

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
