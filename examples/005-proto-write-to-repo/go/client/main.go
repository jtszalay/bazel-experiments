package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	echov1 "github.com/jtszalay/bazel-experiments/examples/proto_write_to_repo/gen/echo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := echov1.NewEchoServiceClient(conn)

	message := "Hello, Echo!"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Echo(ctx, &echov1.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("failed to echo: %v", err)
	}

	fmt.Printf("Echo response: %s\n", resp.GetMessage())
}
