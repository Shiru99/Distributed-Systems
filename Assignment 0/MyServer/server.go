package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	fmt.Printf("Hello, There!\n")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
