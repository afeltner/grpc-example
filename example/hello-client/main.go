// Package main implements a client for hello service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/afeltner/grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

//nolint:gochecknoglobals
var (
	addr = flag.String("addr", "localhost:50537", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

// main is the entry point into the hello-client.  It will do the following:
// * create a grpc connection to the specified address.
// * create a HelloClient for sending requests.
// * call SayHello and print the returned response message.
func main() {
	flag.Parse()

	grpcConnection, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcConnection.Close()

	clientConnection := pb.NewHelloClient(grpcConnection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := clientConnection.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Panicf("could not SayHello: %v", err)
	}
	log.Printf("%s!", response.GetMessage())
}
