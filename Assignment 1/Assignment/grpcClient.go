package main

import (
	"context"
	"fmt"
	"log"
	"time"

	transaction "example.com/transaction"
	"google.golang.org/grpc"
)

/*1. No duplicate Entry */
func PartI() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := transaction.NewBankServiceClient(conn)
	request := &transaction.TranxReq{TransactionID: "1"}

	request1 := &transaction.TranxReq{TransactionID: "2"}
	fmt.Printf("No duplicate Entry: \n")
	resp, _ := client.GetBalance(context.Background(), request)
	fmt.Printf("Receive response => [%v]\n", resp.TranxOutput)

	time.Sleep(1 * time.Second)
	resp_0, _ := client.GetBalance(context.Background(), request1)
	fmt.Printf("Receive response => [%v]\n", resp_0.TranxOutput)
}

/* 2. Implement the logic to generate duplicate transaction */
func PartII() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := transaction.NewBankServiceClient(conn)
	request := &transaction.TranxReq{TransactionID: "1"}

	fmt.Printf("Duplicate Entry: \n")
	resp1, _ := client.GetBalance(context.Background(), request)
	fmt.Printf("Receive response => [%v]\n", resp1.TranxOutput)

	time.Sleep(1 * time.Second)

	resp2, _ := client.GetBalance(context.Background(), request)
	fmt.Printf("Receive response => [%v]\n", resp2.TranxOutput)

	request1 := &transaction.TranxReq{TransactionID: "2"}
	fmt.Printf("No duplicate Entry: \n")
	resp, _ := client.GetBalance(context.Background(), request)
	fmt.Printf("Receive response => [%v]\n", resp.TranxOutput)

	time.Sleep(1 * time.Second)
	resp_0, _ := client.GetBalance(context.Background(), request1)
	fmt.Printf("Receive response => [%v]\n", resp_0.TranxOutput)
}

/* 3. Deposite the amount */
func PartIII() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := transaction.NewBankServiceClient(conn)
	request := &transaction.TranxReq{TransactionID: "1"}

	request2 := &transaction.DepoReq{TransactionID: "3", Amount: "5"}
	fmt.Printf("Deposite Amount: \n")
	resp_3, _ := client.DepositeAmount(context.Background(), request2)
	resp_3_1, _ := client.GetBalance(context.Background(), request)
	fmt.Printf("Receive response => [%v]\n", resp_3.DepoOutput)
	fmt.Printf("Receive response => [%v]\n", resp_3_1.TranxOutput)
}

/* 4. Server Crash */
func PartIV() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := transaction.NewBankServiceClient(conn)
	request := &transaction.TranxReq{TransactionID: "1"}

	request2 := &transaction.DepoReq{TransactionID: "3", Amount: "5"}
	fmt.Printf("Server Crash: \n")
	resp_3, _ := client.DepositeAmount(context.Background(), request2)
	fmt.Printf("Receive response => [%v]\n", resp_3.DepoOutput)

	resp_3_1, _ := client.GetBalance(context.Background(), request)
	fmt.Printf("Receive response => [%v]\n", resp_3_1.TranxOutput)
}

func main() {
	fmt.Printf("Initiating Client...\n")
	PartI()
	PartII()
	PartIII()
	PartIV()
}
