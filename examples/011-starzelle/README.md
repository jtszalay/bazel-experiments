# Starzelle

This example documents the basic setup of using the aspect-gazelle orion extension with with main bazel-gazelle project. Referred to here as "starzelle" because the extension uses starlark.

The orion extension allows you to [extend gazelle](https://github.com/bazel-contrib/bazel-gazelle/blob/master/extend.md) without needing to write your extension in go. This is useful because you can update the logic without needing to rebuild the extension.

The readme for aspect.build's orion extension for gazelle can be found [here](https://github.com/aspect-build/aspect-gazelle/tree/main/language/orion). Their linked public docs require a login but the old docs for the extension can be found in their archived project [here](https://github.com/aspect-build/aspect-cli-legacy/blob/main/docs/starlark.md).

For the purpose of showing how the extension is setup this one simply prints the paths that are visited when you run
```bash
bazel run //:gazelle
```

```
AspectConfigure-print_paths: .bazel
AspectConfigure-print_paths: tools/starzelle
AspectConfigure-print_paths: tools
AspectConfigure-print_paths:
```
