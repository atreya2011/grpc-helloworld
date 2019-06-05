package main

import (
	"context"
	"flag"
	"log"
	"net"

	"gitlab.com/atreya2011/grpc-helloworld/helloworld"
	"google.golang.org/grpc"
)

type server struct{}

var addrFlag = flag.String("addr", ":5000", "Address host:port")

func main() {
	log.Printf("grpc server start on port %v", *addrFlag)
	// Step 1. listen for connections on tcp
	lis, err := net.Listen("tcp", *addrFlag)
	// Always handle errors
	if err != nil {
		log.Fatalf("frak")
	}
	// Create a new grpc server instance
	sv := grpc.NewServer()
	// Register the Greeter Service by passing the new server instance
	// and the server struct created above
	helloworld.RegisterGreeterServer(sv, server{})
	// Serve the listener created above
	sv.Serve(lis)
}

// SayHello The following is the implementation of the SayHello service
// as defined in the proto file. It can be any implemention.
// The below just returns a message "Hi:" with the name
func (s server) SayHello(ctx context.Context, hello *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Message: "Hi " + hello.GetName()}, nil
}
