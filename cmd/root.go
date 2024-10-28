package cmd

import (
	"os"

	"github.com/ShohamBit/TraceeClient/pkg/client"

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
		Use:   "TraceeClient",
		Short: "TraceeClient is a CLI tool for tracee",
		Long:  "Tracee client is the client for tracee api server.",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if connection type is TCP; if not, disable IP and Port flags
			if serverInfo.ConnectionType == client.PROTOCOL_UNIX {
				// Clear IP and Port as they're not used with Unix socket
				serverInfo.IP = client.DefaultIP
				serverInfo.Port = client.DefaultPort
			} else if serverInfo.ConnectionType != client.PROTOCOL_TCP {
				cmd.Println("Invalid connection type. Use 'tcp' or 'unix'.")
				os.Exit(1)
			}
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
	rootCmd.PersistentFlags().StringVar(&serverInfo.UnixSocketPath, "socketPath", client.SOCKET, "Path of the unix socket")
	// TODO: make this flag available only when the connection type is tcp
	rootCmd.PersistentFlags().StringVarP(&serverInfo.IP, "ip", "i", client.DefaultIP, "IP to connect to the remote server")
	rootCmd.PersistentFlags().StringVarP(&serverInfo.Port, "port", "p", client.DefaultPort, "Port to connect to the remote server")

	// Adjust flags based on connection type
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Disable IP and Port flags for unix connections
		if serverInfo.ConnectionType != client.PROTOCOL_TCP {
			cmd.PersistentFlags().Lookup("ip").Hidden = true
			cmd.PersistentFlags().Lookup("port").Hidden = true
		}
	}

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
