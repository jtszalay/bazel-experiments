# Integration Testing Example

This example demonstrates how to write integration tests for your Bazel project using `rules_itest`.
It continues from the [proto-write-to-repo](../005-proto-write-to-repo/README.md) example.

Unlike unit tests that test individual components in isolation, integration tests verify that multiple components work together correctly. In this example, we write integration tests for the echo gRPC service, starting the server as a long-running process and testing it with a real client.

## Prerequisites

- Bazel (for building and running)
- Understanding of Go testing basics

## What This Example Demonstrates

- Using `rules_itest` to define long-running service dependencies
- Starting a gRPC server on a random port for integration tests
- Writing integration tests that interact with real services
- Structuring integration tests separately from unit tests

## Structure

```
006-integration-testing/
├── proto/              # Protocol buffer definitions
├── go/
│   ├── client/         # Echo client
│   ├── server/         # Echo server
│   └── integration/    # Integration tests
│       ├── BUILD.bazel           # itest_service and service_test rules
│       ├── integration_test.go   # Integration test setup
│       └── messages_test.go      # Test cases
```

## Update Bazel Targets

```bash
bazel run //:gazelle
```

## Run Integration Tests

To run the integration tests:

```bash
bazel test //go/integration:integration_test
```

The `rules_itest` framework will:
1. Start the echo server on a random available port
2. Make the server address available to your tests
3. Run your test suite
4. Clean up the server when tests complete

## How It Works

The integration tests use two key rules from `rules_itest`:

- **`itest_service`**: Defines the echo server as a service that should run during tests
- **`service_test`**: Defines the actual test that depends on the running service

The test receives the server address through environment variables and can make real gRPC calls to verify the service behavior.

## Next

Now, move to the [OCI images](../007-hello-oci/README.md) example.
