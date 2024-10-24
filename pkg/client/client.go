package client

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// github.com/ShohamBit/TraceeClient holds the gRPC connection and service client.

type ServerInfo struct {
	IP   string
	Port string
}

func connectToServer(serverInfo ServerInfo) (*grpc.ClientConn, error) {
	//fmt.Printf("Connecting to server on %s...\n", serverInfo.Address())

	//connect to server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(serverInfo.IP+":"+serverInfo.Port, opts...)
	if err != nil {
		log.Fatalf("server is down %v", err)
		return nil, err
	}
	return conn, nil
}
