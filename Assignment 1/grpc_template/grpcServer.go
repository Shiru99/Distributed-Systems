package main

import (
	"context"
	"fmt"
	"./hellopb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {

}

/* // Sample code to read the request and return the response back to client
func (*server) Hello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := request.Name
	response := &hellopb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
} */

func (*server) GetBalance(ctx context.Context, request *transactionpb.TransactionRequest) (*transactionpb.TransactionResponse, error) {
	//Implement the logic to check the transaction already processed if processed add appropriate message into Response
	
	


	//Else Implement the logic to add the transaction id into TransIds array and Trans_Processed.txt file
	
	
	//Implement the logic to read the Balance from Balance.txt
	
	//Wait for 3 seconds
	time.Sleep(3 * time.Second)
	//Implement the logic to add Balance and Transaction ID to the response
	
	return nil
}

func (*server) DepositeAmount(ctx context.Context, request *transactionpb.TransactionRequest) (*transactionpb.TransactionResponse, error) {
	//Implement the logic to check the transaction already processed if processed add appropriate message into Response
	
	


	//Else Implement the logic to add the transaction id into TransIds array and Trans_Processed.txt file


	//Implement the logic to read the balance and add the amount to into balance and write calculated balance back to Balance.txt
	
	
	time.Sleep(3 * time.Second)	

	//Implement the logic add appropriate message and transaction number to response	

	return nil

}

func main() {
	//Read the transaction processed from Trans_Processed.txt into TransIds
	//Sample code to start the server, read the port number from command-line
	/*address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)
	

	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
	*/
}
