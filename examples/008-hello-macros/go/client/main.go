package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	echov1 "github.com/jtszalay/bazel-experiments/examples/hello_macros/gen/echo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "server address")
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := echov1.NewEchoServiceClient(conn)

	message := "Hello, Echo!"
	if len(flag.Args()) > 0 {
		message = flag.Args()[0]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := SendEchoRequest(ctx, client, message)
	if err != nil {
		log.Fatalf("failed to echo: %v", err)
	}

	fmt.Printf("Echo response: %s\n", resp.GetMessage())
}
