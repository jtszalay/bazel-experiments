package main

import (
	"context"
	"net"
	"testing"
	"time"

	echov1 "github.com/jtszalay/bazel-experiments/examples/hello_macros/gen/echo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type mockEchoServer struct {
	echov1.UnimplementedEchoServiceServer
}

func (s *mockEchoServer) Echo(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	return &echov1.EchoResponse{
		Message: req.GetMessage(),
	}, nil
}

func setupTestServer(t *testing.T) (*grpc.ClientConn, func()) {
	buffer := 1024 * 1024
	listener := bufconn.Listen(buffer)

	server := grpc.NewServer()
	echov1.RegisterEchoServiceServer(server, &mockEchoServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			t.Logf("Server error: %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		"passthrough://bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	cleanup := func() {
		conn.Close()
		server.Stop()
		listener.Close()
	}

	return conn, cleanup
}

func TestEchoClient(t *testing.T) {
	conn, cleanup := setupTestServer(t)
	defer cleanup()

	client := echov1.NewEchoServiceClient(conn)

	tests := []struct {
		name    string
		message string
	}{
		{
			name:    "simple echo",
			message: "Hello, Echo!",
		},
		{
			name:    "empty message",
			message: "",
		},
		{
			name:    "long message",
			message: "This is a longer message to test the echo functionality with more content.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := client.Echo(ctx, &echov1.EchoRequest{Message: tt.message})
			if err != nil {
				t.Fatalf("Echo() error = %v", err)
			}

			if resp.GetMessage() != tt.message {
				t.Errorf("Echo() = %v, want %v", resp.GetMessage(), tt.message)
			}
		})
	}
}
