package client

import (
	"TraceeClient/models"
	"context"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
)

// NewTraceeClient initializes a new gRPC client connection.
func NewServiceClient(serverInfo models.ServerInfo) (*serviceClient, error) {
	conn, err := connectToServer(serverInfo)
	if err != nil {
		log.Fatalf("error appear  %v", err)
	}
	//create service client
	return &serviceClient{
		conn:   conn,
		client: pb.NewTraceeServiceClient(conn),
	}, nil
}

// Close the gRPC connection.
func (tc *serviceClient) CloseConnection() {
	if err := tc.conn.Close(); err != nil {
		log.Printf("Failed to close connection: %v", err)
	}
}

/*
if you want to add new options to the client, under this section is where you should add them
*/

// sends a GetVersion request to the server.
func (tc *serviceClient) GetVersion(ctx context.Context, req *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	return tc.client.GetVersion(ctx, req)
}

func (tc *serviceClient) EnableEvent(ctx context.Context, req *pb.EnableEventRequest) (*pb.EnableEventResponse, error) {
	return tc.client.EnableEvent(ctx, req)
}
