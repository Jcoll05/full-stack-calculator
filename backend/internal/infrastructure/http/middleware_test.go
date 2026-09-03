package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCORSMiddleware tests the CORS middleware to ensure it correctly adds CORS headers and handles preflight requests.
func TestCORSMiddleware(t *testing.T) {
	// Create a flag to check if the next handler was called.
	nextCalled := false

	// Create a dummy next handler that sets the nextCalled flag to true.
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	// Wrap the dummy next handler with the CORS middleware.
	handler := corsMiddleware(next)

	// Test normal requests to ensure CORS headers are added and the next handler is called.
	t.Run("adds CORS headers to normal requests", func(t *testing.T) {
		nextCalled = false

		// Create a new HTTP request to the /api/v1/calculate endpoint.
		req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", nil)
		rec := httptest.NewRecorder()

		// Serve the request using the CORS middleware.
		handler.ServeHTTP(rec, req)

		// Check that the CORS headers are present in the response.
		if rec.Header().Get("Access-Control-Allow-Origin") != "http://localhost:5173" {
			t.Errorf("expected CORS origin header")
		}

		// Check that the next handler was called.
		if !nextCalled {
			t.Errorf("expected next handler to be called")
		}
	})

	// Test preflight requests to ensure they are handled correctly and the next handler is not called.
	t.Run("handles preflight requests", func(t *testing.T) {
		nextCalled = false

		req := httptest.NewRequest(http.MethodOptions, "/api/v1/calculate", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		// Check that the response status code is 204 No Content for preflight requests.
		if rec.Code != http.StatusNoContent {
			t.Errorf("expected status 204, got %d", rec.Code)
		}

		// Check that the next handler was not called for preflight requests.
		if nextCalled {
			t.Errorf("expected next handler not to be called")
		}
	})
}