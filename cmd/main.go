package main

import (
	"fmt"
	"github.com/ayubirz/rent-vs-buy-grpc-go/cmd/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9000"
	}

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s with error: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	server.RegisterGrpcServices(grpcServer)

	fmt.Printf("Starting gRPC server on port %s . . .\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s with error: %v", port, err)
	}
}