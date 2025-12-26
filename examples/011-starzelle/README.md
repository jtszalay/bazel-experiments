# Starzelle Example

This example demonstrates how to extend Gazelle with custom Starlark-based extensions using Aspect's Orion.
It continues from the [bazel-query](../010-bazel-query/README.md) example.

"Starzelle" refers to Gazelle extensions written in Starlark (rather than Go). The Aspect Orion extension allows you to [extend Gazelle](https://github.com/bazel-contrib/bazel-gazelle/blob/master/extend.md) without needing to write Go code. This is useful because you can update extension logic without rebuilding Gazelle itself.

## Prerequisites

- Bazel (for building)
- Understanding of Gazelle basics
- Familiarity with Starlark syntax

## What This Example Demonstrates

- Setting up Aspect Orion extension for Gazelle
- Creating custom Gazelle extensions in Starlark
- Configuring `gazelle_binary` with custom languages
- Using environment variables to configure extensions
- Basic Gazelle extension lifecycle (Configure, GenerateRules, etc.)

## Structure

```
011-starzelle/
├── BUILD.bazel              # Custom gazelle_binary with Orion
├── tools/
│   └── starzelle/           # Starlark extension directory
└── proto/
```

## How It Works

The extension is configured in [BUILD.bazel](file:///Users/james/bazel-experiments/examples/011-starzelle/BUILD.bazel):

```starlark
gazelle_binary(
    name = "gazelle_binary",
    languages = DEFAULT_LANGUAGES + [
        "@aspect_gazelle_orion//:orion",
    ],
)

gazelle(
    name = "gazelle",
    env = {
        "ORION_EXTENSIONS_DIR": "tools/starzelle",
    },
    gazelle = ":gazelle_binary",
)
```

## Run Gazelle with Extension

```bash
bazel run //:gazelle
```

This example's extension prints the paths visited during Gazelle execution:

```
AspectConfigure-print_paths: .bazel
AspectConfigure-print_paths: tools/starzelle
AspectConfigure-print_paths: tools
AspectConfigure-print_paths:
```

## Resources

- [Aspect Gazelle Orion Extension](https://github.com/aspect-build/aspect-gazelle/tree/main/language/orion)
- [Archived Starlark Extension Docs](https://github.com/aspect-build/aspect-cli-legacy/blob/main/docs/starlark.md)
- [Extending Gazelle Guide](https://github.com/bazel-contrib/bazel-gazelle/blob/master/extend.md)

## Benefits of Starlark Extensions

- **No compilation needed**: Changes take effect immediately
- **Easier to write**: Starlark is simpler than Go for many users
- **Faster iteration**: No need to rebuild Gazelle binary for changes
- **Good for prototyping**: Test extension ideas quickly

## Limitations

- Less powerful than Go-based extensions
- May have performance overhead for complex operations
- Access to fewer Gazelle internals

## Next

Now, move to the [starzelle OCI](../012-starzelle-oci/README.md) example.
