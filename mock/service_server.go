package mock

import (
	"context"
	"fmt"
	"net"

	"github.com/ShohamBit/TraceeClient/models"
	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"google.golang.org/grpc"
)

var (
	ExpectedVersion string            = "v0.22.0-15-gd09d7fca0d\n" // Match the output format
	serverInfo      models.ServerInfo = models.ServerInfo{IP: models.DefaultIP, Port: models.DefaultPort}
)

// MockServiceServer implements the gRPC server interface for testing
type MockServiceServer struct {
	pb.UnimplementedTraceeServiceServer // Embed the unimplemented server
}

// Start a mock gRPC server
func StartMockServiceServer() (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", serverInfo.IP, serverInfo.Port))
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	pb.RegisterTraceeServiceServer(s, &MockServiceServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			// Handle the error (e.g., log it)
		}
	}()

	return s, nil
}

func (s *MockServiceServer) GetVersion(ctx context.Context, req *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	// Return a mock version response
	return &pb.GetVersionResponse{Version: ExpectedVersion}, nil
}
func (s *MockServiceServer) EnableEvent(ctx context.Context, req *pb.EnableEventRequest) (*pb.EnableEventResponse, error) {
	return &pb.EnableEventResponse{}, nil
}
