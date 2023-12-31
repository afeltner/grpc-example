# grpc-example

This is an example project using grpc calls.  There are 3 parts to it.
* The server that will be listening for requests and return a response.
* A client that makes a request and prints the results.
* The protobuf request/response contract.

## Prerequisites

Install the protocol compiler plugins for Go using the following commands:

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Make sure you PATH contains the GOPATH bin directory.  If not do:

```
export PATH="$PATH:$(go env GOPATH)/bin"
``` 

# To regenerate the proto files do:

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto 
```

# hello-server

This is the main service, listening on a port for requests.

## To build the server

```
cd cmd/hello-server/
go build
```

## To run the server

```
./hello-server -port=50537
```

# hello-client

This is an example client that will send a request to the hello-server.

## To build the client

```
cd example/hello-client
go build
```

## To run the client

```
./hello-client -name=Bob
```
