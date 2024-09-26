package cmd

import (
	"github.com/ShohamBit/TraceeClient/client"
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

	//create service client
	client, err := client.NewServiceClient(serverInfo)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	//get version
	response, err := client.GetVersion(context.Background(), &pb.GetVersionRequest{})
	if err != nil {
		log.Fatalf("Error getting version: %v", err)
	}
	//display version
	fmt.Printf("Version: %+v\n", response.Version)
	client.CloseConnection()
}
