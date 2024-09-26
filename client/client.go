package client

import (
	"github.com/ShohamBit/TraceeClient/models"
	"fmt"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// github.com/ShohamBit/TraceeClient holds the gRPC connection and service client.
type serviceClient struct {
	conn   *grpc.ClientConn
	client pb.TraceeServiceClient
}

// github.com/ShohamBit/TraceeClient holds the gRPC connection and diagnostic client.
type diagnosticClient struct {
	conn   *grpc.ClientConn
	client pb.DiagnosticServiceClient
}

func connectToServer(serverInfo models.ServerInfo) (*grpc.ClientConn, error) {
	fmt.Printf("create a new service client\n")
	fmt.Printf("Connecting to server on %s...\n", serverInfo.Address())

	//connect to server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(serverInfo.Address(), opts...)
	if err != nil {
		log.Fatalf("server is down %v", err)
		return nil, err
	}
	return conn, nil
}
