# Hello Proto Example

This example demonstrates how to use protobuf with `buf` to generate Go bindings for a simple gRPC echo service.

## Structure

```
003-hello-proto/
├── proto/              # Protobuf definitions
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

# Next

Now, move to the [protobuf with gazelle](../004-proto-gazelle/README.md) example.
