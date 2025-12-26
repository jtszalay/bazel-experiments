# Hello Macros Example

This example demonstrates how to create reusable Bazel macros to reduce boilerplate and standardize patterns.
It continues from the [hello-oci](../007-hello-oci/README.md) example.

In the previous example, building OCI images required repeating several rules for each binary (`tar`, `oci_image`, `oci_load`). Macros allow you to encapsulate this pattern into a single reusable function, making your BUILD files cleaner and easier to maintain.

## Prerequisites

- Bazel (for building)
- Docker (for loading and running images)
- Understanding of Bazel rules and targets

## What This Example Demonstrates

- Creating custom Bazel macros in `.bzl` files
- Encapsulating repetitive rule patterns
- Parameterizing macros with attributes
- Organizing macros in a `tools/macros/` directory
- Using macros to simplify BUILD files

## Structure

```
008-hello-macros/
├── tools/
│   └── macros/
│       └── oci/
│           ├── BUILD.bazel
│           └── defs.bzl          # Custom OCI macro definition
├── proto/                        # Protocol buffer definitions
└── go/
    ├── client/                   # Uses the OCI macro
    └── server/                   # Uses the OCI macro
```

## The OCI Macro

Instead of writing multiple rules for each OCI image:

```starlark
# Before: verbose and repetitive
tar(name = "server_app", ...)
oci_image(name = "server_image", ...)
oci_load(name = "server_load", ...)
```

The macro simplifies it to:

```starlark
# After: clean and reusable
oci(
    name = "server_oci",
    target = ":linux_go_cross",
)
```

The macro automatically creates all necessary targets with consistent naming and configuration.

## Update Bazel Targets

```bash
bazel run //:gazelle
```

## Build and Load Images

Load the images into Docker:

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

**Macros vs Rules:**
- **Macros** are evaluated during the loading phase and expand into multiple targets
- **Rules** are evaluated during the analysis phase and create individual targets
- Macros are simpler to write but less powerful than rules

**Benefits of Macros:**
- Reduce code duplication across BUILD files
- Enforce consistent patterns and naming conventions
- Make complex rule chains easier to use
- Simplify maintenance when patterns need to change

## Next

Now, move to the [multirun](../009-multirun-demo/README.md) example.
