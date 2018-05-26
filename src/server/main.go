package main

import (
	"fmt"
	"github.com/calc_service/src/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Start a gRPC server and waits for connection
func main() {
	// Create a listener on TCP port 12345
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 12345))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a server instance
	s := api.Server{}

	// Create a gRPC server object
	grpcServer := grpc.NewServer()

	// Attach the Calc service to the server
	api.RegisterCalcServer(grpcServer, &s)

	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
