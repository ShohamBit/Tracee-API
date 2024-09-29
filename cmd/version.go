package cmd

import (
	"context"

	"github.com/ShohamBit/TraceeClient/client"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display the version of tracee",
	Long:  "this is the version of tracee application you connected to",
	Run: func(cmd *cobra.Command, args []string) {
		displayVersion(cmd, args)
	},
}

func displayVersion(cmd *cobra.Command, _ []string) {

	//create service client
	client, err := client.NewServiceClient(serverInfo)
	if err != nil {
		cmd.PrintErrf("Error creating client: %v", err)
	}
	defer client.CloseConnection()
	//get version
	response, err := client.GetVersion(context.Background(), &pb.GetVersionRequest{})
	if err != nil {
		cmd.PrintErrf("Error getting version: %v", err)
	}
	//display version
	cmd.Print("Version: ", response.Version)
}
