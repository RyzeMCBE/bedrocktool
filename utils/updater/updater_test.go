package updater

import (
	"net/http"
	"testing"
)

func TestIsMissingForkRelease(t *testing.T) {
	tests := []struct {
		name string
		url string
		status int
		want bool
	}{
		{"unpublished fork release", githubLatestReleaseURL, http.StatusNotFound, true},
		{"temporary release API error", githubLatestReleaseURL, http.StatusServiceUnavailable, false},
		{"other URL not found", "https://example.com/resource", http.StatusNotFound, false},
		{"published fork release", githubLatestReleaseURL, http.StatusOK, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMissingForkRelease(tt.url, tt.status); got != tt.want {
				t.Fatalf("isMissingForkRelease(%q, %d) = %t; want %t", tt.url, tt.status, got, tt.want)
			}
		})
	}
}
