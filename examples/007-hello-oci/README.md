# Hello OCI Example

This example demonstrates how to package Go applications as OCI (Open Container Initiative) images using Bazel.
It continues from the [integration-testing](../006-integration-testing/README.md) example.

OCI images (commonly known as Docker images) allow you to package your application with all its dependencies into a standardized container format that can run consistently across different environments.

## Prerequisites

- Bazel (for building)
- Docker (for loading and running images)

## What This Example Demonstrates

- Using `rules_oci` to build container images with Bazel
- Creating minimal container images for Go applications
- Building multi-platform images (linux/amd64, linux/arm64)
- Loading OCI images into Docker
- Understanding OCI image layers and structure

## Structure

```
007-hello-oci/
├── proto/              # Protobuf definitions
└── go/
    ├── client/         # Echo client with OCI image target
    └── server/         # Echo server with OCI image target
```

## Update Bazel Targets

```bash
bazel run //:gazelle
```

## Build OCI Images

Build the OCI images:

```bash
bazel build //go/server:image
bazel build //go/client:image
```

## Load Images into Docker

To load the images into your local Docker daemon:

```bash
bazel run //go/server:load
bazel run //go/client:load
```

This makes the images available to Docker with tags like `bazel-experiments/server:latest` and `bazel-experiments/client:latest`.

## Run with Docker

Start the server:

```bash
docker run -it --rm -p 50051:50051 bazel-experiments/server:latest
```

In another terminal, run the client:

```bash
docker run -it --rm --add-host=host.docker.internal:host-gateway bazel-experiments/client:latest --addr host.docker.internal:50051 "hello"
```

The `--add-host` flag allows the client container to connect to the server running on the host machine.

## Understanding OCI Images

OCI images consist of layers that are stacked on top of each other:
- **Base layer**: Minimal runtime environment (often distroless or scratch for Go)
- **Application layer**: Your compiled binary and any required files
- **Configuration**: Entry point, environment variables, exposed ports

Bazel's `rules_oci` creates efficient images with minimal layers and no unnecessary dependencies.

## Next

Now, move to the [Bazel macros](../008-hello-macros/README.md) example.
