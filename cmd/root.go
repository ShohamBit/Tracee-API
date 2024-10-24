package cmd

import (
	"os"

	"github.com/ShohamBit/TraceeClient/pkg/client"

	"github.com/spf13/cobra"
)

var (
	DefaultIP   = "localhost"
	DefaultPort = "4466"
	TCS         client.ServiceClient    // tracee service client
	TCD         client.DiagnosticClient // tracee diagnostic  client
	serverInfo  client.ServerInfo       = client.ServerInfo{
		IP:   DefaultIP,
		Port: DefaultPort,
	}

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
	rootCmd.AddCommand(disableEventCmd)
	rootCmd.AddCommand(streamEventsCmd)

	//flags
	rootCmd.PersistentFlags().StringVarP(&serverInfo.IP, "ip", "i", DefaultIP, "IP to connect to the remote server")
	rootCmd.PersistentFlags().StringVarP(&serverInfo.Port, "port", "p", DefaultPort, "Port to connect to the remote server")

}

// run root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// expose root command
func GetRootCmd() *cobra.Command {
	return rootCmd
}
