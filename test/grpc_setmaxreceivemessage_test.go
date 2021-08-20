package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/test/bufconn"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var testServer *grpc.Server

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	testServer = s

	pb.RegisterGreeterServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: '%v'", in.GetName())
	return &pb.HelloReply{Message: "Send back " + in.GetName()}, nil
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Test_DefaultReceiveLength(t *testing.T) {

	testServer.SetMaxReceiveMessageSize(4 * 1024 * 1024)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Sending 18 bytes"})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

	// OK if no error
}

func Test_LimitedReceiveLength(t *testing.T) {

	// set server receive to 10 bytes
	testServer.SetMaxReceiveMessageSize(10)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Sending 18 bytes"})
	// OK only if message fails
	if err == nil {
		t.Fatalf("SetMaxReceiveMessageSize failed: %v", err)
	}

	log.Printf("Expected error: %v", err)
	log.Printf("Response: %+v", resp)

}

func Test_DynamicReceiveLength(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// set server receive to default
	testServer.SetMaxReceiveMessageSize(4 * 1024 * 1024)
	log.Print("1. Sending less than MaxReceiveMessageSize")

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Sending less than MaxReceiveMessageSize"})
	if err != nil {
		t.Fatalf("Receive failed: %v", err)
	}
	log.Printf("Response 1: %+v", resp)

	// set server receive to 10 bytes
	log.Print("2. Sending more than MaxReceiveMessageSize")
	testServer.SetMaxReceiveMessageSize(10)

	resp, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Sending more than MaxReceiveMessageSize"})
	// OK only if message fails
	if err == nil {
		t.Fatalf("SetMaxReceiveMessageSize failed: %v", err)
	}

	log.Printf("Expected error: %v", err)
	log.Printf("Response 2(should be <nil>): %+v", resp)

	// set server receive to default
	testServer.SetMaxReceiveMessageSize(4 * 1024 * 1024)
	log.Print("3. Again sending less than MaxReceiveMessageSize")

	resp, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Sending less than MaxReceiveMessageSize"})
	if err != nil {
		t.Fatalf("Receive failed: %v", err)
	}

	log.Printf("Response 3: %+v", resp)
}
