package cmd

import (
	"TraceeClient/models"

	"os"

	"github.com/spf13/cobra"
)

var (
	serverInfo  models.ServerInfo
	defaultIP   = "localhost"
	defaultPort = "4466"

	rootCmd = &cobra.Command{
		Use:   "TraceeClient",
		Short: "TraceeClient is a CLI tool for tracee",
		Long:  `Tracee client is the client for tracee api server.`,
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
	rootCmd.PersistentFlags().StringVarP(&serverInfo.Port, "port", "p", defaultPort, "Port to connect to the remote server")
	rootCmd.PersistentFlags().StringVarP(&serverInfo.IP, "ip", "i", defaultIP, "IP to connect to the remote server")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
