package cmd

import (
	"os"

	"github.com/ShohamBit/traceectl/pkg/client"

	"github.com/spf13/cobra"
)

var (
	TCS        client.ServiceClient    // tracee service client
	TCD        client.DiagnosticClient // tracee diagnostic  client
	serverInfo client.ServerInfo       = client.ServerInfo{
		ConnectionType: client.PROTOCOL_UNIX,
		UnixSocketPath: client.SOCKET,
		IP:             client.DefaultIP,
		Port:           client.DefaultPort,
	}

	rootCmd = &cobra.Command{
		Use:   "trceectl",
		Short: "TraceeCtl is a CLI tool for tracee",
		Long:  "TraceeCtl is the client for the tracee API server.",
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
	rootCmd.PersistentFlags().StringVarP(&serverInfo.ConnectionType, "connectionType", "c", client.PROTOCOL_UNIX, "Connection type (unix|tcp)")
	rootCmd.RegisterFlagCompletionFunc("connectionType", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{client.PROTOCOL_TCP, client.PROTOCOL_UNIX}, cobra.ShellCompDirectiveNoFileComp
	})
	//TODO: add an option to ony use this flag par connection type
	//unix connection type flag
	rootCmd.PersistentFlags().StringVar(&serverInfo.UnixSocketPath, "socketPath", client.SOCKET, "Path of the unix socket")
	//tcp connection type flag
	rootCmd.PersistentFlags().StringVarP(&serverInfo.IP, "ip", "i", client.DefaultIP, "IP to connect to the remote server")
	rootCmd.PersistentFlags().StringVarP(&serverInfo.Port, "port", "p", client.DefaultPort, "Port to connect to the remote server")

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
