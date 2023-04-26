package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "example.com/bob-grpc/pkg/proto/example" // Update the import path to the location of your generated Go protobuf files
)

func main() {
	// Set up the connection to the gRPC server
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// Create a new SayHello client
	client := pb.NewHelloWorldServiceClient(conn)

	// Prepare the SayHelloRequest
	request := &pb.SayHelloRequest{
		Message: "Hello, Server!",
	}

	// Call the SayHello API
	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	// Print the response message
	fmt.Println("Response:", response.Message)
}
