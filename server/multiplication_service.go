package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gRPC_Go/proto"

	"google.golang.org/grpc"
)

// Define the multiplication server
type server struct {
	pb.UnimplementedMultiplicationServiceServer
}

// Implement the multiplication logic
func (s *server) Multiply(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := req.Num1 * req.Num2
	return &pb.CalculationResponse{Result: result}, nil
}

// Main entry point
func main() {
	lis, err := net.Listen("tcp", ":50052") // Listening for multiplication service on port 50052
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMultiplicationServiceServer(grpcServer, &server{}) // Register service with proto alias

	fmt.Println("MultiplicationService is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
