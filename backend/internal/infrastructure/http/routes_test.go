package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Jcoll05/full-stack-calculator/backend/internal/application"
)

// TestNewRouter verifies that the calculator endpoint is correctly registered.
func TestNewRouter(t *testing.T) {
	service := application.NewCalculatorService() // Create the application service.
	handler := NewCalculatorHandler(service)      // Create the HTTP handler.
	router := NewRouter(handler)                  // Create the HTTP router.

	// Test the /api/v1/calculate endpoint.
	t.Run("calculator endpoint", func(t *testing.T) {
		request := httptest.NewRequest(
			http.MethodPost,
			"/api/v1/calculate",
			strings.NewReader(`{"operation":"add","a":10,"b":5}`),
		)

		recorder := httptest.NewRecorder() // Create a ResponseRecorder to capture the response.

		router.ServeHTTP(recorder, request) // Serve the HTTP request using the router.

		// Check if the response status code is 200 OK.
		if recorder.Code != http.StatusOK {
			t.Errorf(
				"NewRouter() status = %d, want %d",
				recorder.Code,
				http.StatusOK,
			)
		}
	})

	// Test an unknown endpoint to ensure it returns a 404 Not Found.
	t.Run("unknown endpoint", func(t *testing.T) {
		request := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/unknown",
			nil,
		)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		// Check if the response status code is 404 Not Found.
		if recorder.Code != http.StatusNotFound {
			t.Errorf(
				"NewRouter() status = %d, want %d",
				recorder.Code,
				http.StatusNotFound,
			)
		}
	})
}
