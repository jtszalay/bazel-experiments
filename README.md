# bazel-experiments

This repository is originally a fork of [bazel-ide](https://github.com/hofbi/bazel-ide).

## Setup

There is a minimal [devcontainer](.devcontainer/devcontainer.json) available providing a minimal set of dependencies such as `git` and [`direnv`](https://direnv.net/) as well as a few configuration steps executed when the container is created.
This set of minimal dependencies and the few configuration steps should be all you need for a local setup outside of the devcontainer.

All other dependencies are managed by Bazel and [`bazel_env.bzl`](https://github.com/buildbuddy-io/bazel_env.bzl).
To get started, run the following command:

```bash
direnv allow .envrc
bazel run //tools:bazel_env
```

Now you should see a list of tools available in your `PATH`.

## Updating tools

`//tools:bazel_env` installs the [multitool](https://github.com/theoremlp/multitool) companion tool which can be used to update the lockfile.

```bash
multitool --lockfile multitool.lock.json update
```

## Examples

- [Hello Bazel](examples/001-hello-bazel/README.md) - Minimal standalone example using Bazel and rules_go with `go_binary`, `go_library`, and `go_test` rules
- [Hello Gazelle](examples/002-hello-gazelle/README.md) - Using Gazelle to automatically generate and update build files for Go programs with external dependencies
- [Hello Proto](examples/003-hello-proto/README.md) - Using protobuf with `buf` to generate Go bindings for a simple gRPC echo service
- [Proto with Gazelle](examples/004-proto-gazelle/README.md) - Using Bazel-managed protobuf with rules_go, generating `.pb.go` files at build time
- [Proto with Write to Repo](examples/005-proto-write-to-repo/README.md) - Writing Bazel-generated protobuf code to the repository for easier IDE integration

## Resources

Some useful resources used for this project:

- [Developer Tooling in Monorepos with bazel_env - feat. Fabian Meumertzheim](https://www.youtube.com/watch?v=TDyUvaXaZrc)
- [Device management: tools on your developers PATH](https://blog.aspect.build/bazel-devenv)
- [Bazel Env](https://github.com/buildbuddy-io/bazel_env.bzl)
- [Bazel 102: Python](https://training.aspect.build/bazel-102)
- [Bazel 104: C++](https://training.aspect.build/bazel-104-c)
