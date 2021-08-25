# gRPC & failure handling for At-Most-Once Semantics

## [GO Installation](https://github.com/6-CSE/HelloWorld/tree/master/Go) :

```
(go1.17) $ wget -c https://dl.google.com/go/go1.17.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
```

To Run Go Program :
```
$ go run HelloWorld.go 
```


## [gRPC Installation](https://grpc.io/docs/languages/go/quickstart/) :

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Project Initialization 
```
$ cd grpc_template/

$ go mod init example.com/m/v2

$ go mod tidy

$ sudo apt  install protobuf-compiler
```

## Project 

* Note : You need to change your option go_package (File: hello.proto) into option go_package = "./;hellopb"; The first param means relative path where the code you want to generate. The path relative to the --go_out , you set in your command.

```
$ protoc --go_out=. --go_opt=paths=source_relative \
         --go-grpc_out=. \
         --go-grpc_opt=paths=source_relative hellopb/hello.proto

$ 
```
