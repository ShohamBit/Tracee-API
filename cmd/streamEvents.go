package cmd

import (
	"github.com/ShohamBit/TraceeClient/client"
	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

var (
	streamEventsCmd = &cobra.Command{
		Use:   "streamEvents",
		Short: "Stream events from tracee",
		Long:  "Stream the events that tracee trace to the client",
		Run: func(cmd *cobra.Command, args []string) {
			streamEvents(cmd, args)
		},
	}
)

func streamEvents(cmd *cobra.Command, args []string) {
	// create service client
	client, err := client.NewServiceClient(serverInfo)
	if err != nil {
		cmd.PrintErrln("Error creating client: ", err)
	}
	defer client.CloseConnection()
	// stream events
	req := &pb.StreamEventsRequest{Policies: args}
	stream, err := client.StreamEvents(cmd.Context(), req)
	if err != nil {
		cmd.PrintErrln("Error calling StreamEvents: ", err)
	}
	// Receive and process streamed responses
	for {
		res, err := stream.Recv()
		if err != nil {
			cmd.PrintErrln("Error receiving streamed event: ", err)
			break
		}
		cmd.Println(res.Event)
	}

}
