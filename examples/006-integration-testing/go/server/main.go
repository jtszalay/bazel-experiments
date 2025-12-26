package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	echov1 "github.com/jtszalay/bazel-experiments/examples/integration_testing/gen/echo/v1"
	"google.golang.org/grpc"
)

type echoServer struct {
	echov1.UnimplementedEchoServiceServer
}

func (s *echoServer) Echo(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	return handleEchoRequest(ctx, req)
}

func main() {
	port := flag.String("port", "50051", "server port")
	flag.Parse()

	lis, err := net.Listen("tcp", "0.0.0.0:"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	echov1.RegisterEchoServiceServer(s, &echoServer{})

	fmt.Printf("Server listening on :%s\n", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
