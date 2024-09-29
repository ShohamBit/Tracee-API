package cmd

import (
	"context"

	"github.com/ShohamBit/TraceeClient/client"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Display Tracee metrics",
	Long:  "This command fetches and displays various metrics from Tracee such as event counts, errors, and lost events.",
	Run: func(cmd *cobra.Command, args []string) {
		displayMetrics(cmd, args)
	},
}

// displayMetrics fetches and prints Tracee metrics
func displayMetrics(cmd *cobra.Command, _ []string) {

	//create service client
	client, err := client.NewDiagnosticClient(serverInfo)
	if err != nil {
		cmd.PrintErrln("Error creating client: ", err)
	}
	defer client.CloseConnection()
	//get metrics
	response, err := client.GetMetrics(context.Background(), &pb.GetMetricsRequest{})
	if err != nil {
		cmd.PrintErrln("Error getting version: ", err)
	}

	// Display the metrics
	cmd.Println("EventCount:", response.EventCount)
	cmd.Println("EventsFiltered:", response.EventsFiltered)
	cmd.Println("NetCapCount:", response.NetCapCount)
	cmd.Println("BPFLogsCount:", response.BPFLogsCount)
	cmd.Println("ErrorCount:", response.ErrorCount)
	cmd.Println("LostEvCount:", response.LostEvCount)
	cmd.Println("LostWrCount:", response.LostWrCount)
	cmd.Println("LostNtCapCount:", response.LostNtCapCount)
	cmd.Println("LostBPFLogsCount:", response.LostBPFLogsCount)
}
