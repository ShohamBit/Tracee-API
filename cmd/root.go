package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

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
)

func init() {

	// commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(metricsCmd)

	//flags
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", defaultPort, "Port to connect to the remote server")
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", defaultIP, "IP to connect to the remote server")

}

func connect() *grpc.ClientConn {

	fmt.Printf("Connecting to server on %s and port %s...\n", ip, port)
	addr := fmt.Sprintf("%s:%s", ip, port)
	//check if server is up
	isServeUp := func(ip string, port string) bool {
		addr := fmt.Sprintf("%s:%s", ip, port)
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
		if err != nil {
			log.Fatalf("server is down %v", err)
			return false // Server is not reachable
		}
		defer conn.Close()
		return true // Server is reachable
	}
	isServeUp(ip, port)
	//connect to server
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No connection established.: %v", err)
	}

	fmt.Printf("Connected to server \n")
	return conn

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
