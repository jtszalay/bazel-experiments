# Bazel Query Example

This example demonstrates how to use `bazel query` to discover and generate build targets dynamically.
It continues from the [multirun-demo](../009-multirun-demo/README.md) example.

Instead of manually maintaining lists of targets (like OCI load targets), you can use `bazel query` to automatically discover them and generate Bazel configuration. This ensures your build configuration stays up-to-date as you add or remove targets.

## Prerequisites

- Bazel (for building and querying)
- Docker (for loading images)
- Basic understanding of shell scripting

## What This Example Demonstrates

- Using `bazel query` to discover targets matching specific patterns
- Automatically generating Bazel configuration files
- Creating scripts that keep configuration synchronized with code
- Using `sh_binary` rules to make scripts runnable via Bazel
- Dynamic target discovery for multirun commands

## Structure

```
010-bazel-query/
├── BUILD.bazel                  # Contains multirun using generated targets
├── images.bzl                   # Generated file with image targets
├── scripts/
│   └── generate_load_images.sh  # Query and generate images.bzl
├── proto/
└── go/
    ├── client/                  # Client with OCI load target
    └── server/                  # Server with OCI load target
```

## How It Works

The [generate_load_images.sh](./scripts/generate_load_images.sh) script:
1. Uses `bazel query` to find all targets ending in `_oci_load`
2. Generates [images.bzl](./images.bzl) with the list of targets
3. The multirun target loads this list and runs all image loads

## Update Bazel Targets

```bash
bazel run //:gazelle
```

## Generate the Image Targets List

Run the query script to regenerate [images.bzl](./images.bzl):

```bash
bazel run //:generate_load_images
```

This discovers all `*_oci_load` targets and writes them to images.bzl:

```starlark
image_targets = [
    "//go/client:client_oci_load",
    "//go/server:server_oci_load",
]
```

## Load All Images

Use the generated multirun target:

```bash
bazel run //:load_images
```

This runs all discovered image load targets in parallel.

## Bazel Query Examples

Find all OCI load targets:
```bash
bazel query 'kind("oci_load", //...)'
```

Find all targets in a package:
```bash
bazel query '//go/server:*'
```

Find all Go binaries:
```bash
bazel query 'kind("go_binary", //...)'
```

## Run with Docker

Start the server:

```bash
docker run -it --rm -p 50051:50051 bazel-experiments/server_oci:latest
```

In another terminal, run the client:

```bash
docker run -it --rm --add-host=host.docker.internal:host-gateway bazel-experiments/client_oci:latest --addr host.docker.internal:50051 "hello"
```

## Benefits

- **Automatic discovery**: No need to manually maintain lists of targets
- **Stays in sync**: Regenerate anytime to pick up new targets
- **Reduces errors**: Eliminates manual list maintenance
- **Scalable**: Works with any number of targets

## Next

Now, move to the [starzelle](../011-starzelle/README.md) example.
