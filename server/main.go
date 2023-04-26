package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	pb "example.com/bob-grpc/pkg/proto/example"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.HelloWorldServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	message := req.GetMessage()
	fmt.Println("Received message:", message)

	// Create the response message
	responseMessage := "Hello from Server!"
	response := &pb.SayHelloResponse{
		Message: responseMessage,
	}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pb.RegisterHelloWorldServiceServer(s, &helloServer{})
	fmt.Println("Server started on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
