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
	defaultPort = "4466"
	conn        *grpc.ClientConn
)

var rootCmd = &cobra.Command{
	Use:   "Tracee-API",
	Short: "This is the client side for tracee",
	Long: `Tracee client is the api for tracee.
	Tracee client can serve you many options of requests you can ask from tracee`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		conn = connect()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// You can add some default logic here if needed
		fmt.Println("Root command executed")
	},
}

func init() {

	// commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(metricsCmd)

	//flags
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", defaultPort, "Port to connect to the remote server")
}

func connect() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Printf("Connecting to server on port %s...\n", port)
	conn, err := grpc.Dial(":"+port, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	fmt.Printf("Connected to server \n")
	return conn
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
