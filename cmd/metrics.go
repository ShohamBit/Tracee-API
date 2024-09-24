package cmd

import (
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
	//create diagnostic grpc client
	client := pb.NewDiagnosticServiceClient(conn)
	metrics, err := client.GetMetrics(context.Background(), &pb.GetMetricsRequest{})
	if err != nil {
		log.Fatalf("Error fetching metrics: %v", err)
	}

	// Display the metrics
	fmt.Println("EventCount:", metrics.EventCount)
	fmt.Println("EventsFiltered:", metrics.EventsFiltered)
	fmt.Println("NetCapCount:", metrics.NetCapCount)
	fmt.Println("BPFLogsCount:", metrics.BPFLogsCount)
	fmt.Println("ErrorCount:", metrics.ErrorCount)
	fmt.Println("LostEvCount:", metrics.LostEvCount)
	fmt.Println("LostWrCount:", metrics.LostWrCount)
	fmt.Println("LostNtCapCount:", metrics.LostNtCapCount)
	fmt.Println("LostBPFLogsCount:", metrics.LostBPFLogsCount)
}
