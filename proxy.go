package main

import (
	"context"
	"log"
	"net/http"

	"github.com/atreya2011/grpc-helloworld/helloworld"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func startGRPCGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := helloworld.RegisterGreeterHandlerFromEndpoint(ctx, mux, *grpcListenAddr, opts)
	if err != nil {
		log.Fatalln(err)
	}

	return http.ListenAndServe(*httpListenAddr, mux)
}
