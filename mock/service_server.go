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
	mockEvents := []*pb.StreamEventsResponse{
		{Event: generateEvent([]string{""})},
		{Event: generateEvent([]string{"policy1"})},
		{Event: generateEvent([]string{"policy2"})},
		{Event: generateEvent([]string{"policy3"})},
		{Event: generateEvent([]string{"policy1", "policy3"})},
		{Event: generateEvent([]string{"policy1", "policy2"})},
		{Event: generateEvent([]string{"policy2", "policy3"})},
		{Event: generateEvent([]string{"policy1", "policy2", "policy3"})},
	}

	// Simulate streaming of events with delays
	for _, event := range mockEvents {
		// If the request has policies, filter the events
		if len(req.Policies) != 0 {
			if hasAnyMatch(req.Policies, event.Event.Policies.Matched) {
				if err := stream.Send(event); err != nil {
					return err
				}
			}
		} else {
			if err := stream.Send(event); err != nil {
				return err
			}
		}
		time.Sleep(100 * time.Millisecond) // Simulate delay between events

	}
	return nil
}

func hasAnyMatch(arr1, arr2 []string) bool {
	// Create a map for the second array to store the elements
	arr2Map := make(map[string]bool)

	// Populate the map with elements from arr2
	for _, val := range arr2 {
		arr2Map[val] = true
	}

	// Check if at least one element in arr1 exists in arr2
	for _, val := range arr1 {
		if arr2Map[val] {
			return true // Return true as soon as we find a match
		}
	}

	// No matches found
	return false
}
func generateEvent(policy []string) *pb.Event {

	return &pb.Event{
		Policies: &pb.Policies{Matched: policy},
	}

}
