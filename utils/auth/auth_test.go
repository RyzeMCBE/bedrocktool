package auth

import (
	"context"
	"strings"
	"testing"
)

func TestLoginWithNoHandlerReturnsError(t *testing.T) {
	a := &authSrv{}
	err := a.Login(context.Background(), nil, "")
	if err == nil {
		t.Fatal("Login should return an error with no auth code handler")
	}
	if !strings.Contains(err.Error(), "authentication code handler") {
		t.Fatalf("unexpected error: %v", err)
	}
}
