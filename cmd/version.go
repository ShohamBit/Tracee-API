package cmd

import (
	"context"

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
	if err := TCS.NewServiceClient(serverInfo); err != nil {
		cmd.PrintErrln("Error creating client: ", err)
	}
	defer TCS.CloseConnection()
	//get version
	response, err := TCS.GetVersion(context.Background(), &pb.GetVersionRequest{})
	if err != nil {
		cmd.PrintErrln("Error getting version: ", err)
	}
	//display version
	cmd.Println("Version: ", response.Version)
}
