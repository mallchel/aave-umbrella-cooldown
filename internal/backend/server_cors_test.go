package backend

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsLocalhostOrigin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		origin string
		want   bool
	}{
		{name: "localhost with port", origin: "http://localhost:9090", want: true},
		{name: "loopback ipv4", origin: "http://127.0.0.1:9090", want: true},
		{name: "loopback ipv6", origin: "http://[::1]:9090", want: true},
		{name: "remote host", origin: "https://example.com", want: false},
		{name: "subdomain localhost", origin: "http://api.localhost.dev", want: false},
		{name: "invalid", origin: "not a url", want: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := isLocalhostOrigin(tc.origin)
			if got != tc.want {
				t.Fatalf("isLocalhostOrigin(%q) = %v, want %v", tc.origin, got, tc.want)
			}
		})
	}
}

func TestWithLocalhostCORS_LocalhostOriginSetsCORSHeaders(t *testing.T) {
	t.Parallel()

	h := withLocalhostCORS(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/withdraw/flows", nil)
	req.Header.Set("Origin", "http://localhost:9090")
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:9090" {
		t.Fatalf("allow-origin = %q, want %q", got, "http://localhost:9090")
	}
	if got := rec.Header().Get("Access-Control-Allow-Methods"); got != "GET, POST, PUT, PATCH, DELETE, OPTIONS" {
		t.Fatalf("allow-methods = %q, want %q", got, "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	}
}

func TestWithLocalhostCORS_RemoteOriginDoesNotSetCORSHeaders(t *testing.T) {
	t.Parallel()

	h := withLocalhostCORS(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/withdraw/flows", nil)
	req.Header.Set("Origin", "https://example.com")
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Fatalf("allow-origin = %q, want empty", got)
	}
	if got := rec.Header().Get("Access-Control-Allow-Methods"); got != "" {
		t.Fatalf("allow-methods = %q, want empty", got)
	}
}
