# Hello Proto Example

This example demonstrates how to use protobuf with `buf` to generate Go bindings for a simple gRPC echo service.

## Structure

```
003-hello-proto/
├── proto/              # Protocol buffer definitions
│   ├── echo.proto      # Echo service definition
│   ├── buf.yaml        # Buf configuration
│   └── buf.gen.yaml    # Buf code generation config
└── go/                 # Go code
    ├── gen/            # Generated protobuf code (created by buf)
    ├── client/         # Echo client
    ├── server/         # Echo server
    └── go.mod
```

## Prerequisites

- `buf` CLI installed (for protobuf code generation)
- Bazel (for building and running)

## Generate Protocol Buffer Code

Before building or running the Go code, you must generate the protobuf bindings:

```bash
cd proto
buf generate
```

This generates Go code in `go/gen/` from the proto definitions.

## Update Dependencies

After generating proto code, tidy the Go module:

```bash
cd go
bazel run @rules_go//go -- mod tidy
```

## Update Bazel Targets

```bash
bazel run //:gazelle
```

## Build and Run

### Run the Server

```bash
bazel run //go/server
```

The server will listen on `:50051`.

### Run the Client

In another terminal:

```bash
bazel run //go/client
```

You can pass a custom message:

```bash
bazel run //go/client -- "Custom message"
```

## Key Learnings

### Protocol Buffer Generation

- **buf.yaml**: Configures linting and breaking change detection for your proto files
- **buf.gen.yaml**: Specifies how to generate code
  - Uses remote plugins from buf.build (no local protoc needed)
  - Outputs to `../go/gen` directory
  - Uses `paths=source_relative` to maintain directory structure

### Go Module Configuration

- **go.mod location**: Placed in `go/` directory, not at the root
- **MODULE.bazel**: Points to `//go:go.mod` instead of `//:go.mod`
- **Import paths**: Use the module path from go.mod (`github.com/jtszalay/bazel-experiments/examples/hello_proto`)
- **Generated code**: Lives in `go/gen` and is imported like any other Go package

### Bazel Integration

- **Go version**: Specified in both `go.mod` (1.25.5) and `MODULE.bazel` (go_sdk.download)
- **rules_go version**: Using 0.59.0 for Go 1.25.5 support
- **Running go commands**: Use `bazel run @rules_go//go -- <command>` to ensure version consistency

### gRPC Server Implementation

- Implements `UnimplementedEchoServiceServer` for forward compatibility
- Uses `grpc.NewServer()` without TLS for local development
- Listens on TCP port 50051

### gRPC Client Implementation

- Uses `grpc.NewClient()` with `insecure.NewCredentials()` for local development
- Sets a 5-second timeout on requests
- Accepts command-line arguments for custom messages
