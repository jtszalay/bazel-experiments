package main

import (
	"context"
	"testing"

	echov1 "github.com/jtszalay/bazel-experiments/examples/proto_write_to_repo/gen/echo/v1"
)

func TestEchoServer_Echo(t *testing.T) {
	s := &echoServer{}

	tests := []struct {
		name    string
		message string
	}{
		{
			name:    "simple message",
			message: "Hello, World!",
		},
		{
			name:    "empty message",
			message: "",
		},
		{
			name:    "unicode message",
			message: "Hello, 世界! 🌍",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &echov1.EchoRequest{
				Message: tt.message,
			}

			resp, err := s.Echo(context.Background(), req)
			if err != nil {
				t.Fatalf("Echo() error = %v", err)
			}

			if resp.GetMessage() != tt.message {
				t.Errorf("Echo() = %v, want %v", resp.GetMessage(), tt.message)
			}
		})
	}
}
