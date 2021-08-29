package main

import (
	"fmt"
	"log"
	"net"

	"example.com/hello"
	"google.golang.org/grpc"
)

func main() {

	fmt.Printf("Initiating Server...\n")
	fmt.Printf("Hello, There!\n")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	myServer := hello.UnimplementedHelloServiceServer{}

	hello.RegisterHelloServiceServer(grpcServer, &myServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
