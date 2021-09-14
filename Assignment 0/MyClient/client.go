package main

import (
	"context"
	"fmt"
	"log"

	hello "example.com/hello"
	"google.golang.org/grpc"
)

func main() {

	fmt.Printf("Initiating Client...\n")
	fmt.Printf("Hello, There!\n")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := hello.NewHelloServiceClient(conn)

	response, err := c.SayGreetings(context.Background(), &hello.Message{MessageCount: 0, Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayGreetings: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
