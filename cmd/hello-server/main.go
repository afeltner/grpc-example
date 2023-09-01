package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/afeltner/grpc-example/proto"
	"google.golang.org/grpc"
)

// server is used to implement HelloServer.
type server struct {
	pb.UnimplementedHelloServer
}

// SayHello implements hello.Hello.
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// main is the entry point into the hello-server service.  It will do the following:
// * verify all program parameters.
// * set up a port to listen on.
// * Create a new grpc server instance and Serve on the listen port.
func main() {
	flag.Parse()

	if errs := verifyProgramArguments(); errs != nil {
		log.Fatal(fmt.Errorf("initialzation failed because: %w", errs))
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	helloServer := grpc.NewServer()
	pb.RegisterHelloServer(helloServer, &server{}) //nolint:exhaustivestruct
	log.Printf("hello-server is listening at %v", listener.Addr())
	if err := helloServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
