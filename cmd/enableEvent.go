package cmd

import (
	"context"

	"github.com/ShohamBit/TraceeClient/client"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

var enableEventCmd = &cobra.Command{
	Use:   "enableEvent [eventNames...]",
	Short: "Enable specified events on the server",
	Long:  "long about the use",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check if args are provided
		if len(args) == 0 {
			cmd.Println("Error: no event names provided. Please specify at least one event to enable.")
			return // Exit if no arguments
		}
		enableEvents(cmd, args)
	},
}

func enableEvents(cmd *cobra.Command, eventNames []string) {
	// Create Tracee gRPC client
	client, err := client.NewServiceClient(serverInfo)
	if err != nil {
		cmd.PrintErrln("Error creating client: ", err)
		return // Exit on error
	}

	// Iterate over event names and enable each one
	for _, eventName := range eventNames {
		_, err := client.EnableEvent(context.Background(), &pb.EnableEventRequest{Name: eventName})
		if err != nil {
			cmd.PrintErrln("Error enabling event:", err)
			continue // Continue on error with the next event
		}
		cmd.Println("Enabled event:", eventName)
	}
}
