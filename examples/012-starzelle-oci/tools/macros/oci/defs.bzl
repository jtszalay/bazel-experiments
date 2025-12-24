load("@rules_go//go:def.bzl", "go_cross_binary")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load")
load("@tar.bzl", "tar")

def _oci_impl(name, target, base, visibility):

    go_cross_binary(
        name = "{}_linux_go_cross".format(name),
        platform = "@rules_go//go/toolchain:linux_amd64",
        target = target,
    )

    tar(
        name = "{}_app".format(name),
        srcs = [":{}_linux_go_cross".format(name)],
        mtree = [
            "usr/local/bin/{} uid=0 gid=0 mode=0755 type=file content=$(execpath {})".format(name, ":{}_linux_go_cross".format(name)),
        ],
    )

    oci_image(
        name = "{}_image".format(name),
        base = base,
        entrypoint = ["/usr/local/bin/{}".format(name)],
        tars = [":{}_app".format(name)],
        visibility = visibility,
    )

    oci_load(
        name = "{}_load".format(name),
        image = ":{}_image".format(name),
        repo_tags = ["bazel-experiments/{}:latest".format(name)],
        visibility = visibility,
    )

oci = macro(
    attrs = {
        "target": attr.label(mandatory = True, configurable = False, doc = "The binary to include in the OCI image"),
        "base": attr.label(default = "@alpine", configurable = False, doc = "The base image to use"),
    },
    implementation = _oci_impl,
)
