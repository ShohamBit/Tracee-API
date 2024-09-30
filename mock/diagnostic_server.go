package mock

import (
	"context"
	"fmt"
	"net"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"google.golang.org/grpc"
)

var (
	ExpectedMetrics pb.GetMetricsResponse = pb.GetMetricsResponse{EventCount: 1, EventsFiltered: 2, NetCapCount: 3,
		BPFLogsCount: 4, ErrorCount: 5, LostEvCount: 6,
		LostWrCount: 7, LostNtCapCount: 8, LostBPFLogsCount: 9}
)

// MockDiagnosticServer implements the gRPC server interface for testing
type MockDiagnosticServer struct {
	pb.UnimplementedDiagnosticServiceServer // Embed the unimplemented server
}

// Start a mock gRPC server
func StartMockDiagnosticServer() (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", serverInfo.IP, serverInfo.Port))
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	pb.RegisterDiagnosticServiceServer(s, &MockDiagnosticServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			// Handle the error (e.g., log it)
		}
	}()

	return s, nil
}

func (s *MockDiagnosticServer) GetMetrics(ctx context.Context, req *pb.GetMetricsRequest) (*pb.GetMetricsResponse, error) {
	return &ExpectedMetrics, nil
}
