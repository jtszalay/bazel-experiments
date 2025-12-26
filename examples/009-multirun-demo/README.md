# Multirun Demo Example

This example demonstrates how to run multiple Bazel targets in parallel or sequence using `rules_multirun`.
It continues from the [hello-macros](../008-hello-macros/README.md) example.

When working with multiple related services or images, you often need to run several Bazel commands (like loading multiple Docker images). `rules_multirun` allows you to define a single target that executes multiple commands, either in parallel for speed or sequentially for ordered execution.

## Prerequisites

- Bazel (for building)
- Docker (for loading images)

## What This Example Demonstrates

- Using `rules_multirun` to execute multiple targets
- Running targets in parallel with `jobs = 0`
- Creating composite commands for common workflows
- Simplifying multi-step operations into single commands

## Structure

```
009-multirun-demo/
├── BUILD.bazel              # Contains multirun target
├── proto/
└── go/
    ├── client/              # Client with OCI image
    └── server/              # Server with OCI image
```

## The Multirun Target

In the root [BUILD.bazel](./BUILD.bazel):

```starlark
multirun(
    name = "load_images",
    commands = [
        "//go/client:client_oci_load",
        "//go/server:server_oci_load",
    ],
    jobs = 0,  # 0 = parallel, positive number = that many jobs
)
```

## Update Bazel Targets

```bash
bazel run //:gazelle
```

## Load All Images with One Command

Instead of running multiple commands:

```bash
# Before: run each command separately
bazel run //go/server:server_oci_load
bazel run //go/client:client_oci_load
```

Use the multirun target:

```bash
# After: run all at once
bazel run //:load_images
```

With `jobs = 0`, both image loads run in parallel, significantly reducing total time.

## Run with Docker

Start the server:

```bash
docker run -it --rm -p 50051:50051 bazel-experiments/server_oci:latest
```

In another terminal, run the client:

```bash
docker run -it --rm --add-host=host.docker.internal:host-gateway bazel-experiments/client_oci:latest --addr host.docker.internal:50051 "hello"
```

## Parallel vs Sequential Execution

- **Parallel (`jobs = 0`)**: Runs all commands simultaneously. Faster but output may be interleaved.
- **Sequential (`jobs = 1`)**: Runs commands one at a time. Slower but output is ordered.
- **Limited parallelism (`jobs = N`)**: Runs up to N commands at once.

## Use Cases

`rules_multirun` is useful for:
- Loading multiple Docker images
- Running multiple formatters or linters
- Starting multiple services for testing
- Any workflow requiring multiple related commands

## Next

Now, move to the [bazel query](../010-bazel-query/README.md) example.
