package main

import (
	"context"
	"fmt"
	"log"
	"net"

	echov1 "github.com/jtszalay/bazel-experiments/examples/bazel_query/gen/echo/v1"
	"google.golang.org/grpc"
)

type echoServer struct {
	echov1.UnimplementedEchoServiceServer
}

func (s *echoServer) Echo(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	return handleEchoRequest(ctx, req)
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	echov1.RegisterEchoServiceServer(s, &echoServer{})

	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
