package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gRPC_Go/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMultiplicationServiceServer
	logs []*pb.LogEntry // Store logs
}

func (s *server) Multiply(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := req.Num1 * req.Num2
	s.logs = append(s.logs, &pb.LogEntry{Operation: "Multiplication", Result: result}) // Log result
	return &pb.CalculationResponse{Result: result}, nil
}

func (s *server) StreamLogs(req *pb.LogRequest, stream pb.MultiplicationService_StreamLogsServer) error {
	for _, logEntry := range s.logs {
		if err := stream.Send(logEntry); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMultiplicationServiceServer(grpcServer, &server{})
	fmt.Println("MultiplicationService is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
