package cmd

import (
	"TraceeClient/client"
	"context"
	"fmt"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Display Tracee metrics",
	Long:  "This command fetches and displays various metrics from Tracee such as event counts, errors, and lost events.",
	Run: func(cmd *cobra.Command, args []string) {
		displayMetrics()
	},
}

// displayMetrics fetches and prints Tracee metrics
func displayMetrics() {

	//create service client
	client, err := client.NewDiagnosticClient(serverInfo)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	//get metrics
	response, err := client.GetMetrics(context.Background(), &pb.GetMetricsRequest{})
	if err != nil {
		log.Fatalf("Error getting version: %v", err)
	}

	// Display the metrics
	fmt.Println("EventCount:", response.EventCount)
	fmt.Println("EventsFiltered:", response.EventsFiltered)
	fmt.Println("NetCapCount:", response.NetCapCount)
	fmt.Println("BPFLogsCount:", response.BPFLogsCount)
	fmt.Println("ErrorCount:", response.ErrorCount)
	fmt.Println("LostEvCount:", response.LostEvCount)
	fmt.Println("LostWrCount:", response.LostWrCount)
	fmt.Println("LostNtCapCount:", response.LostNtCapCount)
	fmt.Println("LostBPFLogsCount:", response.LostBPFLogsCount)
}
