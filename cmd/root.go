package cmd

import (
	"github.com/ShohamBit/TraceeClient/models"

	"os"

	"github.com/spf13/cobra"
)

var (
	serverInfo models.ServerInfo

	rootCmd = &cobra.Command{
		Use:   "TraceeClient",
		Short: "TraceeClient is a CLI tool for tracee",
		Long:  "Tracee client is the client for tracee api server.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {

	// commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(metricsCmd)
	rootCmd.AddCommand(enableEventCmd)

	//flags
	rootCmd.PersistentFlags().StringVarP(&serverInfo.IP, "ip", "i", models.DefaultIP, "IP to connect to the remote server")
	rootCmd.PersistentFlags().StringVarP(&serverInfo.Port, "port", "p", models.DefaultPort, "Port to connect to the remote server")

}

// run root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// expose root command
func NewRootCommand() *cobra.Command {
	return rootCmd
}

func GetServerInfo() models.ServerInfo {
	return serverInfo
}
