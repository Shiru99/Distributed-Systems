# gRPC Basics

## [GO Installation](https://github.com/6-CSE/HelloWorld/tree/master/Go) :

To Run Go Program :
```
$ go run HelloWorld.go 
```


## [gRPC Installation](https://grpc.io/docs/languages/go/quickstart/) :

```
$ sudo apt  install protobuf-compiler

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Project Initialization 
```
$ go mod init example.com

$ go mod tidy
```

## Creating gRPC code using transaction.proto contract

```
$ protoc --go_out=. --go_opt=paths=source_relative \
         --go-grpc_out=. \
         --go-grpc_opt=paths=source_relative transaction/transaction.proto
```


## How To Run

Note : clear Trans_Processed.txt before running grpcServer.go (For each Part I-IV)

```
$ go run grpcServer.go

$ go run grpcClient.go
```

