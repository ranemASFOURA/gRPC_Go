package main

import (
	"context" // Provides context for managing deadlines, cancelation signals, and request-scoped values
	"fmt"     // Provides formatted I/O functions
	"log"     // Provides logging utilities
	"net"     // Provides network-related functionalities

	pb "gRPC_Go/proto" // Import the generated protobuf package

	"google.golang.org/grpc" // Import gRPC library
)

// Define a server struct that implements the MultiplicationService defined in the proto file
type server struct {
	pb.UnimplementedMultiplicationServiceServer                // Embedding the unimplemented server for forward compatibility
	logs                                        []*pb.LogEntry // Slice to store log entries for each multiplication operation
}

// Multiply method: Handles multiplication requests and logs each operation
func (s *server) Multiply(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {

	result := req.Num1 * req.Num2
	// Log the multiplication result by appending a new log entry to the logs slice
	s.logs = append(s.logs, &pb.LogEntry{Operation: "Multiplication", Result: result})

	return &pb.CalculationResponse{Result: result}, nil
}

// StreamLogs method: Streams stored log entries to the client
func (s *server) StreamLogs(req *pb.LogRequest, stream pb.MultiplicationService_StreamLogsServer) error {
	// Iterate through each log entry in the logs slice
	for _, logEntry := range s.logs {
		// Send each log entry to the client
		if err := stream.Send(logEntry); err != nil {
			// If there is an error in streaming, return the error
			return err
		}
	}
	// Return nil when done streaming logs
	return nil
}

// Main function to set up and run the gRPC server
func main() {
	// Start a TCP listener on port 50052 to listen for incoming connections
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err) // Log and exit if there's an error in starting the listener
	}

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()
	// Register the MultiplicationService with the gRPC server
	pb.RegisterMultiplicationServiceServer(grpcServer, &server{})

	fmt.Println("MultiplicationService is running on port 50052...")
	// Start the server and listen for requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err) // Log and exit if there's an error in serving
	}
}
