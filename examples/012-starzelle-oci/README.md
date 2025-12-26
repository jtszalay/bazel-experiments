# Starzelle OCI Example

This example demonstrates how to use a Starlark Gazelle extension to automatically generate OCI build targets.
It continues from the [starzelle](../011-starzelle/README.md) example.

Building on the Starzelle concept, this example shows a practical use case: automatically detecting Go binaries (via `main.go` files) and generating OCI image build targets for them. This eliminates the need to manually add `oci()` macro calls to BUILD files.

## Prerequisites

- Bazel (for building)
- Docker (for loading images)
- Understanding of Gazelle and Starlark extensions

## What This Example Demonstrates

- Creating a practical Gazelle extension that generates build targets
- Using `aspect.SourceGlobs` to discover files
- Automatically generating OCI targets for Go binaries
- Integrating custom Gazelle extensions with existing build rules
- The `# gazelle:oci enabled` directive to control extension behavior

## Structure

```
012-starzelle-oci/
├── BUILD.bazel                  # Custom gazelle_binary with Orion
├── tools/
│   ├── starzelle/
│   │   └── oci.axl              # OCI extension in Starlark
│   └── macros/
│       └── oci/
│           └── defs.bzl         # OCI macro from hello-macros
├── proto/
└── go/
    ├── client/                  # OCI target auto-generated
    └── server/                  # OCI target auto-generated
```

## How the Extension Works

The [oci.axl](file:///Users/james/bazel-experiments/examples/012-starzelle-oci/tools/starzelle/oci.axl) extension:

1. **Searches for `main.go` files** using `aspect.SourceGlobs("**/main.go")`
2. **Generates `oci()` targets** automatically for each directory containing `main.go`
3. **Names targets consistently** as `{dirname}_oci`

The extension is enabled via a Gazelle directive:

```starlark
# gazelle:oci enabled
```

## Run Gazelle

To generate OCI targets automatically:

```bash
bazel run //:gazelle
```

Gazelle will discover all `main.go` files and add corresponding `oci()` macro calls to the BUILD files.

## Build and Load Images

Load all discovered images:

```bash
bazel run //:load_images
```

Or load individual images:

```bash
bazel run //go/server:server_oci_load
bazel run //go/client:client_oci_load
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

## Key Concepts

**Extension Lifecycle:**
1. **`prepare()`**: Defines which files to scan (e.g., `**/main.go`)
2. **`declare()`**: Generates targets based on discovered files

**Benefits:**
- No manual target maintenance for OCI images
- Consistent naming conventions enforced automatically
- New binaries automatically get OCI targets when Gazelle runs
- Combines power of macros with automation of Gazelle

**Compared to Previous Examples:**
- **008-hello-macros**: Manual `oci()` macro calls in BUILD files
- **012-starzelle-oci**: Automatic `oci()` target generation via Gazelle extension

## Next

Now, move to the [GoMock](../013-gomocks_demo/README.md) example.
