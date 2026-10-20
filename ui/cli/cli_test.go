package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLIAuthHandlerPrintsDeviceCode(t *testing.T) {
	var output bytes.Buffer
	handler := newCLIAuthHandler(&output)

	handler.DisplayCode("https://microsoft.com/link", "ABCD-1234")

	for _, part := range []string{"https://microsoft.com/link", "ABCD-1234"} {
		if !strings.Contains(output.String(), part) {
			t.Errorf("Microsoft device login output missing %q: %q", part, output.String())
		}
	}
}

func TestCLIAuthHandlerReportsSuccessfulLogin(t *testing.T) {
	var output bytes.Buffer
	handler := newCLIAuthHandler(&output)
	handler.Successful()

	if !strings.Contains(output.String(), "Authentication successful") {
		t.Fatalf("expected confirmation, got %q", output.String())
	}
}
