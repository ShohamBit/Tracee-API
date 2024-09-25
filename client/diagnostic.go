package client

import (
	"TraceeClient/models"
	"context"
	"log"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
)

// NewTraceeClient initializes a new gRPC client connection.
func NewDiagnosticClient(serverInfo models.ServerInfo) (*diagnosticClient, error) {
	conn, err := connectToServer(serverInfo)
	if err != nil {
		log.Fatalf("error appear  %v", err)
	}
	//create diagnostic client
	return &diagnosticClient{
		conn:   conn,
		client: pb.NewDiagnosticServiceClient(conn),
	}, nil
}

// Close the gRPC connection.
func (tc *diagnosticClient) CloseConnection() {
	if err := tc.conn.Close(); err != nil {
		log.Printf("Failed to close connection: %v", err)
	}
}

/*
if you want to add new options to the client, under this section is where you should add them
*/

// sends a GetMetrics request to the server.
func (tc *diagnosticClient) GetMetrics(ctx context.Context, req *pb.GetMetricsRequest) (*pb.GetMetricsResponse, error) {
	return tc.client.GetMetrics(ctx, req)
}
