package main

import (
	"context"
	"testing"
	"time"

	echov1 "github.com/jtszalay/bazel-experiments/examples/gomocks_demo/gen/echo/v1"
	"go.uber.org/mock/gomock"
)

func TestEchoClientWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockEchoServiceClient(ctrl)

	expectedMessage := "Hello from mock!"
	expectedResponse := &echov1.EchoResponse{
		Message: expectedMessage,
	}

	mockClient.EXPECT().
		Echo(gomock.Any(), &echov1.EchoRequest{Message: expectedMessage}).
		Return(expectedResponse, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := mockClient.Echo(ctx, &echov1.EchoRequest{Message: expectedMessage})
	if err != nil {
		t.Fatalf("Echo() error = %v", err)
	}

	if resp.GetMessage() != expectedMessage {
		t.Errorf("Echo() = %v, want %v", resp.GetMessage(), expectedMessage)
	}
}
