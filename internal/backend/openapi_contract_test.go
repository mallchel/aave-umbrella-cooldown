package backend

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type recordingServer struct {
	Unimplemented
	withdrawFlowsCalled bool
}

func (s *recordingServer) ListWithdrawFlows(w http.ResponseWriter, _ *http.Request, _ ListWithdrawFlowsParams) {
	s.withdrawFlowsCalled = true
	w.WriteHeader(http.StatusNoContent)
}

func TestGeneratedRoutesAreRegistered(t *testing.T) {
	handler := Handler(Unimplemented{})

	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{name: "chart", path: "/chart", wantStatus: http.StatusNotImplemented},
		{name: "daily series data", path: "/daily-series-data", wantStatus: http.StatusNotImplemented},
		{name: "health", path: "/healthz", wantStatus: http.StatusNotImplemented},
		{name: "withdraw flows", path: "/withdraw-flows", wantStatus: http.StatusNotImplemented},
		{name: "unknown route", path: "/not-found", wantStatus: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, tt.path, nil)

			handler.ServeHTTP(recorder, request)

			if recorder.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", recorder.Code, tt.wantStatus)
			}
		})
	}
}

func TestGeneratedWithdrawFlowsParameterValidation(t *testing.T) {
	server := &recordingServer{}
	handler := HandlerWithOptions(server, ChiServerOptions{
		ErrorHandlerFunc: func(w http.ResponseWriter, _ *http.Request, err error) {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		},
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/withdraw-flows?limit=not-an-int", nil)

	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
	}
	if server.withdrawFlowsCalled {
		t.Fatal("withdraw flows handler was called after invalid generated parameter binding")
	}
	if !strings.Contains(recorder.Body.String(), "limit") {
		t.Fatalf("response body = %q, want it to mention limit", recorder.Body.String())
	}
}

func TestWithdrawFlowsContractValidation(t *testing.T) {
	server := NewServer(nil)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/withdraw-flows?event_type=invalid", nil)

	server.routes().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusBadRequest)
	}
	if !strings.Contains(recorder.Body.String(), "event_type must be one of: request, withdraw") {
		t.Fatalf("response body = %q, want event_type validation error", recorder.Body.String())
	}
}
