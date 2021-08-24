package main

import (
	"context"
	"fmt"
	"./hellopb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	//Sample code to connect to server

	/* opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := hellopb.NewHelloServiceClient(cc)
	request := &hellopb.HelloRequest{Name: "IIT Dharwad"}

	//Sample code to call server service
	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)
	*/
	
	time.Sleep(1 * time.Second)
	
	//Implement the logic to generate duplicate transaction


	
	//Wait for the transanctions complete
	time.Sleep(6 * time.Second)

	//Display the response
	/*log.Printf("Response1: Transaction Number - %v : %v", response.TransId, response.Data)
	log.Printf("Response2: Transaction Number - %v : %v",response2.TransId, response2.Data)
	*/
	

	//Implement rest of the transactions
}
