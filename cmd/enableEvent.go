package cmd

import (
	"github.com/ShohamBit/TraceeClient/client"
	"context"
	"fmt"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

var enableEventCmd = &cobra.Command{
	Use:   "enableEvent [eventNames...]",
	Short: "Enable specified events on the server",
	Long:  "long about the use",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		enableEvents(args)
	},
}

func enableEvents(eventNames []string) {
	// create Tracee grpc client
	client, err := client.NewServiceClient(serverInfo)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	for _, eventName := range eventNames {
		_, err := client.EnableEvent(context.Background(), &pb.EnableEventRequest{Name: eventName})
		if err != nil {
			log.Fatalf("Error enabling event: %v", err)
		}
		fmt.Printf("Enabled event: %s\n", eventName)
	}
}
