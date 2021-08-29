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
$ go mod init example.com/hello-grpc

$ go mod tidy
```

<!-- ## Project 

* Note : You need to change your option go_package (File: hello.proto) into option go_package = "./;hello"; The first param means relative path where the code you want to generate. The path relative to the --go_out , you set in your command.

```
$ protoc --go_out=. --go_opt=paths=source_relative \
         --go-grpc_out=. \
         --go-grpc_opt=paths=source_relative hello/hello.proto

$ 
``` -->
