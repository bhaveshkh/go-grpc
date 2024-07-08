# go-grpc

This repository contains the implementation of gRPC services in Go.

## Building the Proto File

To compile the proto file, use the following command:

```sh
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

## Steps to Run the Code
### Running the Server
First, navigate to the server directory and run all the Go files:

```sh
go run ./main.go ./client_stream.go ./server_stream.go ./bi_stream.go ./unary.go
```

Running the Client
Next, navigate to the client directory and run all the Go files:

```sh
go run ./main.go ./client_stream.go ./server_stream.go ./bi_stream.go ./unary.go
```

This will start both the server and client, allowing you to test the gRPC services.
