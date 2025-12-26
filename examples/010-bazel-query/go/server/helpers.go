package main

import (
	"context"
	"log"

	echov1 "github.com/jtszalay/bazel-experiments/examples/bazel_query/gen/echo/v1"
)

func handleEchoRequest(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	log.Printf("Received: %s", req.GetMessage())
	return &echov1.EchoResponse{
		Message: req.GetMessage(),
	}, nil
}
