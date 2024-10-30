package cmd

import (
	pb "github.com/aquasecurity/tracee/api/v1beta1"

	"github.com/ShohamBit/traceectl/pkg/cmd/formatter"
	"github.com/spf13/cobra"
)

var formatFlag string
var outputFlag string
var streamCmd = &cobra.Command{
	Use:   "stream [policies...]",
	Short: "Stream events from tracee",
	Long: `Stream Management:
  	- traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  	- traceectl stream describe <stream_name>
  	- traceectl stream list
  	- traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  	- traceectl stream delete <stream_name>
  	- traceectl stream connect <stream_name>
  	- traceectl stream set-default <stream_name>
  	- traceectl stream pause <stream_name>
  	- traceectl stream resume <stream_name>
	`,
	Run: func(cmd *cobra.Command, args []string) {

		//stream(cmd, args)
	},
}

func init() {
	//subcommands
	streamCmd.AddCommand(createStreamCmd)
	streamCmd.AddCommand(describeStreamCmd)
	streamCmd.AddCommand(listStreamCmd)
	streamCmd.AddCommand(updateStreamCmd)
	streamCmd.AddCommand(deleteStreamCmd)
	streamCmd.AddCommand(connectStreamCmd)
	streamCmd.AddCommand(setDefaultStreamCmd)
	streamCmd.AddCommand(pauseStreamCmd)
	streamCmd.AddCommand(resumeStreamCmd)
	//stream events flags
	streamCmd.Flags().StringVarP(&formatFlag, "format", "f", "json", "Output format (json[default]|table|template) currently only support table")
	// only support stdout for now
	streamCmd.Flags().StringVarP(&outputFlag, "output", "o", "stdout", "Output destination ")
}

var createStreamCmd = &cobra.Command{
	Use:   "create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]",
	Short: "Create a new stream",
	Long:  `Creates a new event stream with a specified name.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var describeStreamCmd = &cobra.Command{
	Use:   "describe <stream_name>",
	Short: "describe a stream",
	Long:  `Retrieves the details of a specific stream by its name.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var listStreamCmd = &cobra.Command{
	Use:   "list",
	Short: "list streams",
	Long:  `Lists all available streams, providing a brief summary of each.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var updateStreamCmd = &cobra.Command{
	Use:   "update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]",
	Short: "update a stream",
	Long:  `Updates a stream by its name.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var deleteStreamCmd = &cobra.Command{
	Use:   "delete <stream_name>",
	Short: "delete a stream",
	Long:  `Deletes a stream by its name.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var connectStreamCmd = &cobra.Command{
	Use:   "connect <stream_name>",
	Short: "connect a stream",
	Long:  `Connects a stream by its name.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var setDefaultStreamCmd = &cobra.Command{
	Use:   "set-default <stream_name>",
	Short: "set default stream",
	Long:  `Sets a stream as the default stream.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var pauseStreamCmd = &cobra.Command{
	Use:   "pause <stream_name>",
	Short: "pause a stream",
	Long:  `Pauses a stream by its name.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var resumeStreamCmd = &cobra.Command{
	Use:   "resume <stream_name>",
	Short: "resume a stream",
	Long:  `Resumes a stream by its name.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func stream(cmd *cobra.Command, args []string) {
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
		cmd.PrintErrln("Error calling Stream: ", err)
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
