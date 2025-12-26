package main

import (
	"context"

	echov1 "github.com/jtszalay/bazel-experiments/examples/hello_proto/gen"
)

func SendEchoRequest(ctx context.Context, client echov1.EchoServiceClient, message string) (*echov1.EchoResponse, error) {
	return client.Echo(ctx, &echov1.EchoRequest{Message: message})
}
