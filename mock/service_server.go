package mock

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/ShohamBit/TraceeClient/models"
	pb "github.com/aquasecurity/tracee/api/v1beta1"
	"google.golang.org/grpc"
)

var (
	ExpectedVersion string            = "v0.22.0-15-gd09d7fca0d" // Match the output format
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
func (s *MockServiceServer) DisableEvent(ctx context.Context, req *pb.DisableEventRequest) (*pb.DisableEventResponse, error) {
	return &pb.DisableEventResponse{}, nil
}

/*
\stream events
*/

// StreamEvents simulates the server-side streaming RPC
// there are 3 policies loaded in the mock server
// policy1, policy2 and policy3
// and the server will return 8 events in total
func (s *MockServiceServer) StreamEvents(req *pb.StreamEventsRequest, stream pb.TraceeService_StreamEventsServer) error {
	// Define mock events to send
	//this create events from policies that the client request for
	// which means no need to check if you need to stream something
	mockEvents := CreateEventsFromPolicies(req.Policies)

	// Simulate streaming of events with delays
	for _, event := range mockEvents {
		if err := stream.Send(event); err != nil {
			return err
		}

	}
	time.Sleep(100 * time.Millisecond) // Simulate delay between events
	return nil
}

func generateEvent(policy []string) *pb.Event {
	return &pb.Event{
		Policies: &pb.Policies{Matched: policy},
	}

}

// this creates events from policies that the client request for
func CreateEventsFromPolicies(policy []string) []*pb.StreamEventsResponse {
	if len(policy) <= 1 {
		return []*pb.StreamEventsResponse{
			{Event: generateEvent([]string{""})},
		}
	} else {
		// this nee+ds to be work on
		return []*pb.StreamEventsResponse{
			{Event: generateEvent([]string{"policy1"})},
			{Event: generateEvent([]string{"policy2"})},
			{Event: generateEvent([]string{"policy3"})},
			{Event: generateEvent([]string{"policy1", "policy2"})},
			{Event: generateEvent([]string{"policy1", "policy3"})},
			{Event: generateEvent([]string{"policy2", "policy3"})},
			{Event: generateEvent([]string{"policy1", "policy2", "policy3"})},
		}
	}
}
