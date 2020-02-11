package main

import (
	"context"
	"log"

	"github.com/atreya2011/grpc-helloworld/helloworld"

	"google.golang.org/grpc"
)

func client() {
	// Connect with the grpc server listening on port 5000 using an insecure connection
	// This creates a new connection
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	// Handle the error as usual
	if err != nil {
		log.Fatalln(err)
	}
	// Close the connection
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	res, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "atreya", Age: "35", DobYear: 1984})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Response : %v\n", res.GetMessage())
}
