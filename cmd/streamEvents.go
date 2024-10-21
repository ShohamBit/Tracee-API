package cmd

import (
	"github.com/ShohamBit/TraceeClient/client"
	"github.com/ShohamBit/TraceeClient/cmd/formatter"
	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

var formatFlag string
var outputFlag string

var (
	streamEventsCmd = &cobra.Command{
		Use:   "streamEvents [policy1] [policy2]...",
		Short: "Stream events from tracee",
		Long:  "Stream the events that tracee trace",
		Run: func(cmd *cobra.Command, args []string) {
			streamEvents(cmd, args)
		},
	}
)

func init() {
	streamEventsCmd.Flags().StringVarP(&formatFlag, "format", "f", "table", "Output format (json|table|template[default]) currently only support table")
	// only support stdout for now
	streamEventsCmd.Flags().StringVarP(&outputFlag, "output", "o", "stdout", "Output destination ")
}
func streamEvents(cmd *cobra.Command, args []string) {
	// Create service client
	client, err := client.NewServiceClient(serverInfo)
	if err != nil {
		cmd.PrintErrln("Error creating client: ", err)
	}
	defer client.CloseConnection()

	// create stream from client
	req := &pb.StreamEventsRequest{Policies: args}
	stream, err := client.StreamEvents(cmd.Context(), req)
	if err != nil {
		cmd.PrintErrln("Error calling StreamEvents: ", err)
	}

	//add check for the output flag
	//TODO:support only table format for now
	switch formatFlag {
	case "table":
		tableStreamEvents(cmd, args, stream)
	case "json":
		jsonStreamEvents(cmd, args, stream)
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
	tbl := formatter.InitTable(outputFlag)
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

		// Print each event as a row in the table
		formatter.PrintTableRow(tbl, res.Event)
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
