package cmd

import (
	pb "github.com/aquasecurity/tracee/api/v1beta1"

	"github.com/ShohamBit/TraceeClient/pkg/cmd/formatter"
	"github.com/spf13/cobra"
)

var formatFlag string
var outputFlag string
var streamEventsCmd = &cobra.Command{
	Use:   "streamEvents [policies...]",
	Short: "Stream events from tracee",
	Long:  "long about the use",
	Run: func(cmd *cobra.Command, args []string) {
		streamEvents(cmd, args)
	},
}

func init() {

	//stream events flags
	streamEventsCmd.Flags().StringVarP(&formatFlag, "format", "f", "json", "Output format (json[default]|table|template) currently only support table")
	// only support stdout for now
	streamEventsCmd.Flags().StringVarP(&outputFlag, "output", "o", "stdout", "Output destination ")
}

func streamEvents(cmd *cobra.Command, args []string) {
	// Create service client
	err := TCS.NewServiceClient(serverInfo)
	if err != nil {
		cmd.PrintErrln("Error creating client: ", err)
	}
	defer TCS.CloseConnection()

	// create stream from client
	req := &pb.StreamEventsRequest{Policies: args}
	stream, err := TCS.StreamEvents(cmd.Context(), req)
	if err != nil {
		cmd.PrintErrln("Error calling StreamEvents: ", err)
	}

	//add check for the output flag
	//TODO:support only table and json format for now
	switch formatFlag {
	case "json":
		jsonStreamEvents(cmd, args, stream)
	case "table":
		tableStreamEvents(cmd, args, stream)
	case "template": // go template
		fallthrough
	default:
		cmd.PrintErrln("Error: output format not supported")
		return
	}
}

// tableStreamEvents prints events in a table format
func tableStreamEvents(cmd *cobra.Command, _ []string, stream pb.TraceeService_StreamEventsClient) {
	// Init table header before streaming starts
	tbl := formatter.New(formatFlag, outputFlag, cmd)
	tbl.PrintTableHeaders()
	// Receive and process streamed responses
	for {
		res, err := stream.Recv()
		if err != nil {
			// Handle the error that occurs when the server closes the stream
			if err.Error() == "EOF" {
				break
			}
			cmd.PrintErrln("Error receiving streamed event: ", err)
		}
		tbl.PrintTableRow(res.Event)

	}
}

// jsonStreamEvents prints events in json format
func jsonStreamEvents(cmd *cobra.Command, _ []string, stream pb.TraceeService_StreamEventsClient) {
	// Receive and process streamed responses
	for {
		res, err := stream.Recv()
		if err != nil {
			// Handle the error that occurs when the server closes the stream
			if err.Error() == "EOF" {
				break
			}
			cmd.PrintErrln("Error receiving streamed event: ", err)
		}
		// Print each event as a row in json format
		formatter.PrintJSON(cmd, res.Event, outputFlag)
	}
}
