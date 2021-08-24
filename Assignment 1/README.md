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