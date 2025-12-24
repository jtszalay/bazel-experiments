# Hello OCI

## TODO

* Discuss OCI
    * Layers

It continues from the hello-oci [example](../007-hello-oci/README.md).

This example requires that `docker` be installed on the host. `docker` is not managed by bazel in this repo.

```bash
bazel run //go/server:load
bazel run //go/client:load
```

If on a mac use the following:
```bash
bazel run --config=linux_arm64 //go/server:load
bazel run --config=linux_arm64 //go/client:load
```

```bash
docker run -it --rm -p 50051:50051 bazel-experiments/server_oci:latest
```

```bash
docker run -it --rm --add-host=host.docker.internal:host-gateway bazel-experiments/client_oci:latest --addr host.docker.internal:50051 "hello"
```

# Next
