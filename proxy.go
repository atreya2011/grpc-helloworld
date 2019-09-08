package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/atreya2011/grpc-helloworld/helloworld"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var helloEndPoint = flag.String("hello", "localhost:5000", "endpoint of GreeterService")

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := helloworld.RegisterGreeterHandlerFromEndpoint(ctx, mux, *helloEndPoint, opts)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Listening...")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatalln(err)
	}
}
