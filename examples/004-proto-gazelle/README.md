# Proto with gazelle Example

This example demonstrates how to use bazel managed protobuf with rules_go.
It continues from the [hello-proto](../003-hello-proto/README.md) example.

For this example we have a similar structure to the previous example except rather than
generating `.pb.go` files with buf to a package in the `go` tree we use bazel to provide those at build time.

## Structure

```
004-proto-gazelle/
├── proto/              # Protobuf definitions
│   ├── echo.proto      # Echo service definition
└── go/                 # Go code
    ├── client/         # Echo client
    ├── server/         # Echo server
    └── go.mod
```

## Prerequisites

- Bazel (for building and running)

## Update Bazel Targets

```bash
bazel run //:gazelle
```

Doing so updates the build files to ensure the dependency tree is up to date for go -> proto.

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

Now, either move to the [integration testing](../006-integration-testing/README.md) example.
Or, take a detour to the [write to repo](../005-proto-write-to-repo/README.md) example to learn how we could write the bazel generated proto files to the repo.
