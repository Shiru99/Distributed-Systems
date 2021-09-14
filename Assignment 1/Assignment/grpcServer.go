package main

import (
	"context"
	"fmt"
	"time"

	"bufio"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	transaction "example.com/transaction"
	"google.golang.org/grpc"
)

type server struct {
	transaction.UnimplementedBankServiceServer
}

var TransIds []string
var index int
var balance int

type Listener int

type Response struct {
	Data    string
	TransNo string
}

type Request struct {
	Data    int
	TransNo string
}

//function to check if the array contains the given element
func contains(array []string, toFind string) bool {
	for _, item := range array {
		if item == toFind {
			return true
		}
	}
	return false
}

/* // Sample code to read the request and return the response back to client
func (*server) Hello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := request.Name
	response := &hellopb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
} */

func (*server) GetBalance(ctx context.Context, request *transaction.TranxReq) (*transaction.TranxRes, error) {
	//Implement the logic to check the transaction already processed if processed add appropriate message into Response
	TransactionID := request.TransactionID
	// fmt.Printf("TransactionID: %v", TransactionID)

	found := contains(TransIds, TransactionID)
	if found {
		response := &transaction.TranxRes{
			TranxOutput: "Trasaction Id already exists",
		}

		return response, nil
	}

	//Else Implement the logic to add the transaction id into TransIds array and Trans_Processed.txt file
	TransIds = append(TransIds, TransactionID)
	index++

	//adding the transaction to Trans_Processed file
	f, err := os.OpenFile("Trans_Processed.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString("\n" + request.TransactionID) //args.TransNo)
	if err != nil {
		fmt.Println(err)
		defer f.Close()
	}

	//Implement the logic to read the Balance from Balance.txt
	data, err := ioutil.ReadFile("Balance.txt")
	if err != nil {
		fmt.Errorf("Balance file reading error %s", err)
	}

	//Wait for 3 seconds
	// time.Sleep(3 * time.Second)
	//Implement the logic to add Balance and Transaction ID to the response
	response := &transaction.TranxRes{
		TranxOutput: string(data),
	}

	return response, nil
}

func (*server) DepositeAmount(ctx context.Context, request *transaction.DepoReq) (*transaction.DepoRes, error) {
	//Implement the logic to check the transaction already processed if processed add appropriate message into Response
	TransactionID := request.TransactionID
	// fmt.Printf("TransactionID: %v", TransactionID)

	found := contains(TransIds, TransactionID)
	if found {
		response := &transaction.DepoRes{
			DepoOutput: "Trasaction Id already exists",
		}

		return response, nil
	}

	//Else Implement the logic to add the transaction id into TransIds array and Trans_Processed.txt file

	//adding the transaction no to TransId array
	TransIds = append(TransIds, TransactionID)
	index++

	//adding the transaction to Trans_Processed file
	f, err := os.OpenFile("Trans_Processed.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	_, err = f.WriteString("\n" + TransactionID)
	if err != nil {
		fmt.Println(err)
		defer f.Close()
	}

	//Implement the logic to read the balance and add the amount to into balance and write calculated balance back to Balance.txt

	data, err := ioutil.ReadFile("Balance.txt")
	if err != nil {
		fmt.Println("Balance file reading error", err)
	}

	balance, _ := strconv.Atoi(string(data))
	offset, _ := strconv.Atoi(request.Amount)
	newBalance := balance + offset

	//write calculated balance back to Balance.txt
	file, err := os.OpenFile("Balance.txt", os.O_RDWR, 0644)
	if err != nil {
		return nil, nil
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(newBalance))
	if err != nil {
		return nil, nil
	}

	//Implement the logic to add Balance and Transaction ID to the response
	// response.Data = "Your Amount is deposited into the account successfully."

	// return nil

	time.Sleep(3 * time.Second)

	//Implement the logic add appropriate message and transaction number to response
	response := &transaction.DepoRes{
		DepoOutput: "Your Amount is deposited into the account successfully.",
	}

	return response, nil
}

func main() {

	fmt.Printf("Initiating Server...\n")

	//Read the transaction processed from Trans_Processed.txt into TransIds
	file, err := os.Open("Trans_Processed.txt")

	if err != nil {
		log.Fatalf("Failed to open Trans_Processed file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		TransIds = append(TransIds, scanner.Text())
		index++
	}

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Server is listening...")

	s := grpc.NewServer()
	transaction.RegisterBankServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
