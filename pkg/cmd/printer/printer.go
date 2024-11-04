package printer

import (
	"github.com/ShohamBit/traceectl/pkg/cmd/formatter"
	pb "github.com/aquasecurity/tracee/api/v1beta1"
)

func StreamEvents(format *formatter.Formatter, args []string, stream pb.TraceeService_StreamEventsClient) {

	//add check for the output flag
	//TODO:support only table and json format for now
	switch format.Format {
	case formatter.FormatJSON:
		jsonStreamEvents(args, stream, format)
	case formatter.FormatTable:
		tableStreamEvents(args, stream, format)
	case formatter.FormatGoTpl: // gotemplate
		fallthrough
	default:
		format.CMD.PrintErrln("Error: output format not supported")
		return
	}
}

// tableStreamEvents prints events in a table format
func tableStreamEvents(_ []string, stream pb.TraceeService_StreamEventsClient, tbl *formatter.Formatter) {
	// Init table header before streaming starts
	tbl.PrintTableHeaders()
	// Receive and process streamed responses
	for {
		res, err := stream.Recv()
		if err != nil {
			// Handle the error that occurs when the server closes the stream
			if err.Error() == "EOF" {
				break
			}
			tbl.CMD.PrintErrln("Error receiving streamed event: ", err)
		}
		tbl.PrintTableRow(res.Event)

	}
}

// jsonStreamEvents prints events in json format
func jsonStreamEvents(_ []string, stream pb.TraceeService_StreamEventsClient, tbl *formatter.Formatter) { // Receive and process streamed responses
	for {
		res, err := stream.Recv()
		if err != nil {
			// Handle the error that occurs when the server closes the stream
			if err.Error() == "EOF" {
				break
			}
			tbl.CMD.PrintErrln("Error receiving streamed event: ", err)
		}
		// Print each event as a row in json format
		tbl.PrintJSON(res.Event)
	}
}
