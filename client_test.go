package main

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/atreya2011/grpc-helloworld/helloworld"
	hwmock "github.com/atreya2011/grpc-helloworld/mock_helloworld"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func TestSayHello(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterClient(conn)
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "atreya", Age: "35", DobYear: 1984})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	want := "Hi atreya 35 1984"
	t.Logf("Response: %+v", resp)
	if resp.GetMessage() != want {
		t.Fatalf("got %s, want %s", resp.GetMessage(), want)
	}
}

func TestSayHelloAgain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := hwmock.NewMockGreeterClient(ctrl)
	mockGreeterClient.EXPECT().SayHello(
		gomock.Any(),
		gomock.Any(),
	).Return(&helloworld.HelloResponse{Message: "Mocked Interface"}, nil)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := mockGreeterClient.SayHello(ctx, &helloworld.HelloRequest{Name: "unit_test"})
	if err != nil || r.Message != "Mocked Interface" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Message)
}
