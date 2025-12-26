# GoMocks Demo Example

This example demonstrates how to use GoMock with Bazel to generate mock implementations for testing gRPC services.
It continues from the [starzelle-oci](../012-starzelle-oci/README.md) example.

Mock objects allow you to test code in isolation by replacing external dependencies (like gRPC clients) with controllable test doubles. This example shows how to integrate `bazel_gomock` to automatically generate mock implementations of gRPC interfaces for use in unit tests.

## Prerequisites

- Bazel (for building and testing)
- Understanding of Go interfaces and testing
- Familiarity with mocking concepts

## What This Example Demonstrates

- Using `bazel_gomock` to generate mock implementations
- Testing gRPC clients with mocks instead of real servers
- Configuring `gomock` rules in BUILD files
- Writing tests with mock expectations
- Integrating generated mocks into test targets

## Structure

```
013-gomocks_demo/
├── BUILD.bazel              # Includes bazel_gomock dependency
├── proto/                   # Echo service definition
└── go/
    ├── client/
    │   ├── BUILD.bazel      # gomock rule and test setup
    │   ├── main.go
    │   ├── client_test.go       # Regular integration test
    │   └── client_mock_test.go  # Mock-based unit test
    └── server/
        ├── main.go
        └── server_test.go       # Unit test of server logic
```

## How GoMock Works

The `gomock` rule generates mock implementations:

```starlark
gomock(
    name = "mock_echo_service",
    out = "mock_echo_service.go",
    interfaces = ["EchoServiceClient"],
    library = "//proto:echov1_go_proto",
    package = "main",
)
```

This generates a `NewMockEchoServiceClient()` function that can be used in tests.

## Update Bazel Targets

```bash
bazel run //:gazelle
```

## Run Tests

Run all tests including mock-based tests:

```bash
bazel test //go/client:client_test
bazel test //go/server:server_test
```

## Understanding the Mock Test

The [client_mock_test.go](file:///Users/james/bazel-experiments/examples/013-gomocks_demo/go/client/client_mock_test.go) shows how to use mocks:

```go
// Create a mock controller
ctrl := gomock.NewController(t)
defer ctrl.Finish()

// Create mock client
mockClient := NewMockEchoServiceClient(ctrl)

// Set expectations
mockClient.EXPECT().
    Echo(gomock.Any(), &echov1.EchoRequest{Message: expectedMessage}).
    Return(expectedResponse, nil)

// Use the mock in your test
resp, err := mockClient.Echo(ctx, &echov1.EchoRequest{Message: expectedMessage})
```

## Benefits of Mocking

- **Faster tests**: No need to start real servers
- **Isolated testing**: Test client logic without server dependencies
- **Controlled behavior**: Simulate errors, timeouts, and edge cases
- **Deterministic**: No flaky tests due to network issues
- **Better coverage**: Test error paths that are hard to trigger with real services

## Mock vs Integration Tests

This example includes both:

- **`client_mock_test.go`**: Fast unit test using mocks (tests client logic in isolation)
- **`client_test.go`**: Integration test with real server (tests full interaction)

## Build and Run the Application

The application itself works the same as previous examples:

```bash
bazel run //:load_images
docker run -it --rm -p 50051:50051 bazel-experiments/server_oci:latest
docker run -it --rm --add-host=host.docker.internal:host-gateway bazel-experiments/client_oci:latest --addr host.docker.internal:50051 "hello"
```

## Key Concepts

**GoMock Integration:**
1. Add `bazel_gomock` dependency to MODULE.bazel
2. Use `gomock()` rule to generate mocks from interfaces
3. Include generated mocks in test `srcs`
4. Import `go.uber.org/mock/gomock` in tests

**Best Practices:**
- Generate mocks for interfaces, not concrete types
- Use mocks for external dependencies (databases, APIs, gRPC clients)
- Keep integration tests for end-to-end validation
- Use `gomock.Any()` for arguments you don't care about
- Always call `ctrl.Finish()` to verify all expectations were met

## Resources

- [GoMock Documentation](https://github.com/uber-go/mock)
- [bazel_gomock Rules](https://github.com/bazel-contrib/bazel_gomock)
