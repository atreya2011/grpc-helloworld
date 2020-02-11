package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/atreya2011/grpc-helloworld/helloworld"
	"google.golang.org/grpc"
)

type server struct{}

var grpcListenAddr = flag.String("grpc", ":5000", "grpc listen address")
var httpListenAddr = flag.String("http", ":8080", "http listen address")

func main() {
	flag.Parse()
	log.Printf("grpc server start on port %v\n", *grpcListenAddr)
	// Step 1. listen for connections on tcp
	lis, err := net.Listen("tcp", *grpcListenAddr)
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
	go sv.Serve(lis)
	log.Printf("grpc gateway start on port %v", *httpListenAddr)
	if err := startGRPCGateway(); err != nil {
		log.Fatalln(err)
	}
}

// SayHello The following is the implementation of the SayHello service
// as defined in the proto file. It can be any implemention.
// The below just returns a message "Hi:" with the name
func (s server) SayHello(ctx context.Context, hello *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	msg := fmt.Sprintf("%s %s %d", hello.GetName(), hello.GetAge(), hello.GetDobYear())
	return &helloworld.HelloResponse{Message: "Hi " + msg}, nil
}
