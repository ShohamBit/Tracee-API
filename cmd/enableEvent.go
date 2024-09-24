package cmd

import (
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
	client := pb.NewTraceeServiceClient(conn)

	for _, eventName := range eventNames {
		enableEvent(client, eventName)
	}
}
func enableEvent(client pb.TraceeServiceClient, eventName string) {
	_, err := client.EnableEvent(context.Background(), &pb.EnableEventRequest{Name: eventName})
	if err != nil {
		log.Fatalf("Error enabling event: %v", err)
	}
	fmt.Printf("Enabled event %s\n", eventName)
}
