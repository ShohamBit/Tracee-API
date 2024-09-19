package cmd

import (
	"context"
	"fmt"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display the version of tracee",
	Long:  "this is the version of tracee application you connected to",
	Run: func(cmd *cobra.Command, args []string) {
		displayVersion()
	},
}

func displayVersion() {
	if conn == nil {
		log.Fatal("No connection established. Please run the root command to connect first.")
	}

	client := pb.NewTraceeServiceClient(conn)
	response, err := client.GetVersion(context.Background(), &pb.GetVersionRequest{})
	if err != nil {
		log.Fatalf("Error getting version: %v", err)
	}

	fmt.Printf("Version: %+v\n", response.Version)
}
