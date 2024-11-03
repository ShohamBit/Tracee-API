package mock

import (
	"context"

	pb "github.com/aquasecurity/tracee/api/v1beta1"
)

func (s *MockServiceServer) EnableEvent(ctx context.Context, req *pb.EnableEventRequest) (*pb.EnableEventResponse, error) {
	return &pb.EnableEventResponse{}, nil
}
func (s *MockServiceServer) DisableEvent(ctx context.Context, req *pb.DisableEventRequest) (*pb.DisableEventResponse, error) {
	return &pb.DisableEventResponse{}, nil
}
