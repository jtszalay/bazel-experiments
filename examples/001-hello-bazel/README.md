# rules_go "Hello World" example

This example can originally be found in the [`rules_go` examples](https://github.com/bazel-contrib/rules_go/blob/master/examples/hello/README.md). This readme includes additional information and exercises for the readers.

This directory contains a minimal, standalone example using Bazel and rules_go. It shows how to use the [`go_binary`](https://github.com/bazel-contrib/rules_go/blob/master/docs/go/core/rules.md#go_binary), [`go_library`](https://github.com/bazel-contrib/rules_go/blob/master/docs/go/core/rules.md#go_library), and [`go_test`](https://github.com/bazel-contrib/rules_go/blob/master/docs/go/core/rules.md#go_test) rules without requiring any other dependencies.

To run the binary:

```bash
bazel run //:hello
```

To test the library:

```bash
bazel test //:hello_test
```

For an explanation and an introduction to Bazel, see [Bazel Tutorial: Build a Go Project](https://bazel.build/start/go).

<!-- End original README.md -->

# Exercises

## Exercise 1: Caching

After building and running the tests, try running them again.
You should see
    ```
    (cached) PASSED in 0.0s
    ```
bazel has cached the first run of the tests.

Changing either of the go files will cause the tests to be re-run.

## Exercise 2: Adding a new file

Add a new `hello_lib.go` file to the project and move the `sayHello` function into it.
Now try `bazel run //:hello` to see the new function in action.
Bazel will encounter an error
    ```
    hello.go:8:2: undefined: sayHello
    ```
because our new `hello_lib.go` is not known to via the `go_library` rule.

If we update the `go_library` rule to include our new file, Bazel will be able to find it and build the binary successfully.

Updating build files can be tedious, but it is a necessary step to ensure that Bazel knows about all the files in your project.

We can use another tool to help us with this. [gazelle](https://github.com/bazelbuild/bazel-gazelle) is a tool that can automatically generate and update build files based on the source code in your project.

# Next

Now, move to the [gazelle](../002-hello-gazelle/README.md) example.
