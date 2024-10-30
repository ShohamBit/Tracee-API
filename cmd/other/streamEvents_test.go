package cmd_test

// import (
// 	"bytes"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/ShohamBit/traceectl/mock"
// 	"github.com/ShohamBit/traceectl/models"
// 	pb "github.com/aquasecurity/tracee/api/v1beta1"
// )

// var LoadedPolicies = []string{"policy1", "policy2", "policy3"}
// var StreamEventTests = []models.TestCase{
// 	// TODO: write test for json formats
// 	{
// 		Name:           "No Policies",
// 		Args:           []string{"streamEvents"}, // Removed the empty string argument
// 		ExpectedOutput: mock.CreateEventsFromPolicies([]string{}),
// 	},
// 	{
// 		Name:           "Single policy",
// 		Args:           []string{"streamEvents", LoadedPolicies[0]},
// 		ExpectedOutput: mock.CreateEventsFromPolicies([]string{LoadedPolicies[0]}),
// 	},
// 	{
// 		Name:           "Two policies",
// 		Args:           []string{"streamEvents", LoadedPolicies[0], LoadedPolicies[1]},
// 		ExpectedOutput: mock.CreateEventsFromPolicies(LoadedPolicies[0:2]),
// 	},
// 	{
// 		Name:           "Three policies",
// 		Args:           []string{"streamEvents", LoadedPolicies[0], LoadedPolicies[1], LoadedPolicies[2]},
// 		ExpectedOutput: mock.CreateEventsFromPolicies(LoadedPolicies),
// 	},
// }

// func TestStreamEvent(t *testing.T) {
// 	for _, test := range StreamEventTests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			// Start the mock server
// 			mockServer, err := mock.StartMockServiceServer()
// 			if err != nil {
// 				t.Fatalf("Failed to start mock server: %v", err)
// 			}
// 			defer mockServer.Stop() // Ensure the server is stopped after the test

// 			// Wait for the server to start
// 			time.Sleep(100 * time.Millisecond)

// 			// Capture output
// 			var buf bytes.Buffer
// 			rootCmd.SetOut(&buf)
// 			rootCmd.SetErr(&buf)

// 			// Set arguments for the test
// 			rootCmd.SetArgs(test.Args)

// 			// Execute the command
// 			if err := rootCmd.Execute(); err != nil {
// 				t.Fatalf("Execute() failed: %v", err)
// 			}

// 			// Get the expected output
// 			if expectedEvents, ok := test.ExpectedOutput.([]*pb.StreamEventsResponse); ok {
// 				// Split the actual output by newlines
// 				actualEvents := strings.Split(strings.TrimSpace(buf.String()), "\n")

// 				// Check if the number of events match
// 				if len(actualEvents) != len(expectedEvents) {
// 					t.Errorf("Expected %d events, got %d", len(expectedEvents), len(actualEvents))
// 					return
// 				}

// 				// Compare each event
// 				for i, expected := range expectedEvents {
// 					if actualEvents[i] != expected.Event.String() {
// 						t.Errorf("Expected event %d: %q\nGot: %q", i, expected.Event.String(), actualEvents[i])
// 					}
// 				}
// 			} else {
// 				t.Errorf("Type assertion failed, expected output is not []*pb.StreamEventsResponse")
// 			}
// 		})
// 	}
// }
