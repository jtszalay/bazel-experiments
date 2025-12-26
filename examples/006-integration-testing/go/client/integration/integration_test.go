package integration_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	client "github.com/jtszalay/bazel-experiments/examples/integration_testing/client"
	echov1 "github.com/jtszalay/bazel-experiments/examples/integration_testing/gen/echo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestEchoServerIntegration(t *testing.T) {
	port := os.Getenv("ECHO_SERVER_PORT")
	if port == "" {
		t.Fatal("ECHO_SERVER_PORT environment variable not set")
	}

	addr := fmt.Sprintf("localhost:%s", port)
	t.Logf("Connecting to echo server at %s", addr)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	echoClient := echov1.NewEchoServiceClient(conn)

	tests := []struct {
		name    string
		message string
	}{
		{"simple message", "Hello, World!"},
		{"empty message", ""},
		{"unicode message", "Hello, 世界! 🌍"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqCtx, reqCancel := context.WithTimeout(ctx, 5*time.Second)
			defer reqCancel()

			resp, err := client.SendEchoRequest(reqCtx, echoClient, tt.message)
			if err != nil {
				t.Fatalf("Echo failed: %v", err)
			}

			if resp.GetMessage() != tt.message {
				t.Errorf("Echo response mismatch: got %q, want %q", resp.GetMessage(), tt.message)
			}
		})
	}
}
