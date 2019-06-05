package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"gitlab.com/atreya2011/grpc-helloworld/helloworld"

	"google.golang.org/grpc"
)

var addrFlag = flag.String("addr", "localhost:5000", "server address host:post")

func main() {
	// Connect with the grpc server listening on port 5000 using an insecure connection
	// This creates a new connection
	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())
	// Handle the error as usual
	if err != nil {
		log.Fatalln(err)
	}
	// Close the connection
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	res, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "atreya"})
	fmt.Printf("Response : %v\n", res.GetMessage())
}
