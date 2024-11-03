package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ShohamBit/traceectl/pkg/client"
	pb "github.com/aquasecurity/tracee/api/v1beta1"

	"github.com/spf13/cobra"
)

var (
	TCS        client.ServiceClient    // tracee service client
	TCD        client.DiagnosticClient // tracee diagnostic  client
	serverInfo client.ServerInfo       = client.ServerInfo{
		ConnectionType: client.PROTOCOL_UNIX,
		UnixSocketPath: client.SOCKET,
		ADDR:           client.DefaultIP + ":" + client.DefaultPort,
	}

	rootCmd = &cobra.Command{
		Use:   "trceectl [flags] [options]",
		Short: "TraceeCtl is a CLI tool for tracee",
		Long:  "TraceeCtl is the client for the tracee API server.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {

	// commands
	rootCmd.AddCommand(streamCmd)
	rootCmd.AddCommand(eventCmd)
	rootCmd.AddCommand(pluginCmd)
	rootCmd.AddCommand(policyCmd)

	//other commends
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(metricsCmd)
	rootCmd.AddCommand(diagnoseCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(versionCmd)

	//flags
	rootCmd.PersistentFlags().StringVarP(&serverInfo.ConnectionType, "connectionType", "c", client.PROTOCOL_UNIX, "Connection type (unix|tcp)")
	rootCmd.RegisterFlagCompletionFunc("connectionType", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{client.PROTOCOL_TCP, client.PROTOCOL_UNIX}, cobra.ShellCompDirectiveNoFileComp
	})
	//TODO: add an option to ony use this flag par connection type
	//unix connection type flag
	rootCmd.PersistentFlags().StringVar(&serverInfo.UnixSocketPath, "socketPath", client.SOCKET, "Path of the unix socket")
	//tcp connection type flag
	rootCmd.PersistentFlags().StringVarP(&serverInfo.ADDR, "server", "s", client.DefaultIP+":"+client.DefaultPort, "The address and port of the Kubernetes API server")

}

var connectCmd = &cobra.Command{
	Use:   "connect [<stream_name>]",
	Short: "Connect to the server",
	Long:  "Connects to a stream and displays events in real time.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var metricsCmd = &cobra.Command{
	Use:   "metrics [--output <format>]",
	Short: "Display Tracee metrics",
	Long:  "Retrieves metrics about Tracee's performance and resource usage.",
	Run: func(cmd *cobra.Command, args []string) {
		displayMetrics(cmd, args)
	},
}
var diagnoseCmd = &cobra.Command{
	Use:   "diagnose [--component <component_name>]",
	Short: "Collect diagnostic information to help troubleshoot issues",
	Long:  "Collects diagnostic information to help troubleshoot issues.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var logsCmd = &cobra.Command{
	Use:   "logs [--filter <filter>]",
	Short: "Display log messages from Tracee",
	Long:  "Displays log messages from Tracee, optionally filtered.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows the status of the Tracee Daemon and its components",
	Long:  "Shows the status of the Tracee Daemon and its components.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var configCmd = &cobra.Command{
	Use:   "config [set|get|update] [<option>=<value>] [--file <config_file>]",
	Short: "View or modify the Tracee Daemon configuration at runtime.",
	Long:  `View or modify the Tracee Daemon configuration at runtime.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display the version of tracee",
	Long:  "this is the version of tracee application you connected to",
	Run: func(cmd *cobra.Command, args []string) {
		displayVersion(cmd, args)
	},
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

// displayMetrics fetches and prints Tracee metrics
func displayMetrics(cmd *cobra.Command, _ []string) {

	//create service client
	if err := TCD.NewDiagnosticClient(serverInfo); err == nil {
		cmd.PrintErrln("Error creating client: ", err)
	}
	defer TCD.CloseConnection()
	//get metrics
	response, err := TCD.GetMetrics(context.Background(), &pb.GetMetricsRequest{})
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
