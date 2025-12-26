# Proto Write to Repo Example

This example demonstrates how to write bazel-generated protobuf code to the repository for easier IDE integration and debugging.
It continues from the proto-gazelle [example](../004-proto-gazelle/README.md).

Unlike the previous example where `.pb.go` files are generated at build time in bazel's cache, this example writes them to the source tree under `go/gen/`. This provides out of the box IDE support since tools can see the generated code directly.

## Prerequisites

- Bazel (for building and running)

## What This Example Demonstrates

- Writing Bazel-generated protobuf code to the source tree
- Creating custom `write_go_proto_srcs` rules
- Configuring Gazelle to resolve imports to written-to-source targets
- Improving IDE integration by making generated code visible
- Trade-offs between build-time generation vs checked-in generated code

## Structure

```
005-proto-write-to-repo/
├── proto/              # Protobuf definitions
│   ├── echo.proto      # Echo service definition
├── go/                 # Go code
│   ├── gen/            # Generated protobuf code (written to source)
│   │   └── echo/v1/    # Generated echo service code
│   ├── client/         # Echo client
│   └── server/         # Echo server
└── bazel/              # Custom Bazel rules
    └── write_go_generated_srcs.bzl  # Rule for writing generated code
```

## Generate Protocol Buffer Code

To write the generated protobuf code to the repository:

```bash
bazel run //go/gen/echo/v1:write_generated_protos
```

This writes the `.pb.go` files to `go/gen/echo/v1/` in your source tree.

## Update Bazel Targets

```bash
bazel run //:gazelle
```

This updates the build files to ensure the dependency tree is up to date.

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

## Key Differences from Previous Example

- Generated `.pb.go` files are committed to the repository under `go/gen/`
- The `go_library` targets reference local source files instead of bazel-generated outputs
- Custom `write_go_proto_srcs` rule synchronizes generated code to source tree
- Easier IDE integration since generated code is visible in the workspace
- Requires gazelle configuration to resolve imports to written-to-source targets instead of `go_proto_library` targets

The downside to this approach is you must manage the write to repo rules and regenerate the sources when they change.

## Next

Now, move to the [integration testing](../006-integration-testing/README.md) example.
