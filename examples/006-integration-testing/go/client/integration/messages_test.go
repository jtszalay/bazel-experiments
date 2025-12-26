package integration_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	client "github.com/jtszalay/bazel-experiments/examples/integration_testing/client"
	echov1 "github.com/jtszalay/bazel-experiments/examples/integration_testing/gen/echo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestDefaultMessage(t *testing.T) {
	port := os.Getenv("ECHO_SERVER_PORT")
	if port == "" {
		t.Fatal("ECHO_SERVER_PORT environment variable not set")
	}

	addr := fmt.Sprintf("localhost:%s", port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	echoClient := echov1.NewEchoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test with default empty message
	resp, err := client.SendEchoRequest(ctx, echoClient, "")
	if err != nil {
		t.Fatalf("Echo with default message failed: %v", err)
	}

	if resp.GetMessage() != "" {
		t.Errorf("Expected empty default message, got %q", resp.GetMessage())
	}
}

func TestVariousMessages(t *testing.T) {
	port := os.Getenv("ECHO_SERVER_PORT")
	if port == "" {
		t.Fatal("ECHO_SERVER_PORT environment variable not set")
	}

	addr := fmt.Sprintf("localhost:%s", port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	echoClient := echov1.NewEchoServiceClient(conn)

	testCases := []struct {
		name     string
		message  string
		validate func(t *testing.T, input, output string)
	}{
		{
			name:    "short message",
			message: "Hi",
			validate: func(t *testing.T, input, output string) {
				if input != output {
					t.Errorf("got %q, want %q", output, input)
				}
			},
		},
		{
			name:    "long message",
			message: strings.Repeat("abcdefgh", 100),
			validate: func(t *testing.T, input, output string) {
				if input != output {
					t.Errorf("long message not echoed correctly")
				}
			},
		},
		{
			name:    "special characters",
			message: "!@#$%^&*()_+-=[]{}|;':\",./<>?",
			validate: func(t *testing.T, input, output string) {
				if input != output {
					t.Errorf("got %q, want %q", output, input)
				}
			},
		},
		{
			name:    "multiline message",
			message: "Line 1\nLine 2\nLine 3",
			validate: func(t *testing.T, input, output string) {
				if input != output {
					t.Errorf("got %q, want %q", output, input)
				}
			},
		},
		{
			name:    "numeric message",
			message: "123456789",
			validate: func(t *testing.T, input, output string) {
				if input != output {
					t.Errorf("got %q, want %q", output, input)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := client.SendEchoRequest(ctx, echoClient, tc.message)
			if err != nil {
				t.Fatalf("Echo failed for %q: %v", tc.name, err)
			}

			tc.validate(t, tc.message, resp.GetMessage())
		})
	}
}

func TestConcurrentRequests(t *testing.T) {
	port := os.Getenv("ECHO_SERVER_PORT")
	if port == "" {
		t.Fatal("ECHO_SERVER_PORT environment variable not set")
	}

	addr := fmt.Sprintf("localhost:%s", port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	echoClient := echov1.NewEchoServiceClient(conn)

	// Send multiple concurrent requests
	const numRequests = 10
	results := make(chan error, numRequests)

	for i := 0; i < numRequests; i++ {
		go func(id int) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			message := fmt.Sprintf("Concurrent message %d", id)
			resp, err := client.SendEchoRequest(ctx, echoClient, message)
			if err != nil {
				results <- fmt.Errorf("request %d failed: %w", id, err)
				return
			}

			if resp.GetMessage() != message {
				results <- fmt.Errorf("request %d: got %q, want %q", id, resp.GetMessage(), message)
				return
			}

			results <- nil
		}(i)
	}

	// Collect results
	for i := 0; i < numRequests; i++ {
		if err := <-results; err != nil {
			t.Error(err)
		}
	}
}
