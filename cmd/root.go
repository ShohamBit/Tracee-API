package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port        string
	ip          string
	defaultIP   = "localhost"
	defaultPort = "4466"
	conn        *grpc.ClientConn

	rootCmd = &cobra.Command{
		Use:   "TraceeClient",
		Short: "TraceeClient is a CLI tool for tracee",
		Long:  `Tracee client is the client for tracee api server.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Don't run the connect function for the root command itself
			if cmd == cmd.Root() {
				return
			}

			// For all other subcommands, run the connect function
			conn = connect()
		},
	}
)

func init() {

	// commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(metricsCmd)
	rootCmd.AddCommand(enableEventCmd)

	//flags
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", defaultPort, "Port to connect to the remote server")
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", defaultIP, "IP to connect to the remote server")

}

func connect() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	var addr string = ip + ":" + port
	fmt.Printf("Connecting to server on %s \n", addr)
	//check if server is up
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatalf("server is down %v", err)
	}
	fmt.Printf("Connected to server \n")
	return conn

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
