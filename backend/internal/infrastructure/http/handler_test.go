package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Jcoll05/full-stack-calculator/backend/internal/application"
	"github.com/Jcoll05/full-stack-calculator/backend/internal/domain"
)

// mockCalculator is a test implementation of the Calculator interface.
type mockCalculator struct {
	result float64
	err    error
}

// Calculate returns the configured result and error.
func (m *mockCalculator) Calculate(operation domain.Operation, a, b float64) (float64, error) {
	return m.result, m.err
}

// TestCalculatorHandler_InternalServerError verifies that unexpected service errors
// are returned as HTTP 500 responses.
func TestCalculatorHandler_InternalServerError(t *testing.T) {
	service := &mockCalculator{
		err: errors.New("unexpected error"),
	}

	handler := NewCalculatorHandler(service)

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/calculate",
		strings.NewReader(`{"operation":"add","a":10,"b":5}`),
	)

	recorder := httptest.NewRecorder()

	handler.Calculate(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf(
			"Calculate() status = %d, want %d",
			recorder.Code,
			http.StatusInternalServerError,
		)
	}

	var response errorResponse

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}

	if response.Error != "internal server error" {
		t.Errorf(
			"Calculate() error = %q, want %q",
			response.Error,
			"internal server error",
		)
	}
}

// TestCalculatorHandler_Calculate tests the Calculate method of the CalculatorHandler.
func TestCalculatorHandler_Calculate(t *testing.T) {
	service := application.NewCalculatorService()
	handler := NewCalculatorHandler(service)

	tests := []struct {
		name       string
		method     string
		body       string
		wantStatus int
		wantResult *float64
		wantError  string
		wantAllow  string
	}{
		{
			name:       "addition",
			method:     http.MethodPost,
			body:       `{"operation":"add","a":10,"b":5}`,
			wantStatus: http.StatusOK,
			wantResult: float64Ptr(15),
		},
		{
			name:       "division",
			method:     http.MethodPost,
			body:       `{"operation":"divide","a":10,"b":2}`,
			wantStatus: http.StatusOK,
			wantResult: float64Ptr(5),
		},
		{
			name:       "square root without b",
			method:     http.MethodPost,
			body:       `{"operation":"sqrt","a":25}`,
			wantStatus: http.StatusOK,
			wantResult: float64Ptr(5),
		},
		{
			name:       "division by zero",
			method:     http.MethodPost,
			body:       `{"operation":"divide","a":10,"b":0}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "cannot divide by zero",
		},
		{
			name:       "addition with zero operand",
			method:     http.MethodPost,
			body:       `{"operation":"add","a":10,"b":0}`,
			wantStatus: http.StatusOK,
			wantResult: float64Ptr(10),
		},
		{
			name:       "percentage",
			method:     http.MethodPost,
			body:       `{"operation":"percentage","a":25}`,
			wantStatus: http.StatusOK,
			wantResult: float64Ptr(0.25),
		},
		{
			name:       "power",
			method:     http.MethodPost,
			body:       `{"operation":"power","a":2,"b":3}`,
			wantStatus: http.StatusOK,
			wantResult: float64Ptr(8),
		},
		{
			name:       "negative square root",
			method:     http.MethodPost,
			body:       `{"operation":"sqrt","a":-25}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "cannot calculate square root of a negative number",
		},
		{
			name:       "unsupported operation",
			method:     http.MethodPost,
			body:       `{"operation":"invalid","a":10,"b":5}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "unsupported operation",
		},
		{
			name:       "invalid JSON",
			method:     http.MethodPost,
			body:       `{"operation":"add","a":10,`,
			wantStatus: http.StatusBadRequest,
			wantError:  "invalid request body",
		},
		{
			name:       "missing operation",
			method:     http.MethodPost,
			body:       `{"a":10,"b":5}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "operation is required",
		},
		{
			name:       "missing operand a",
			method:     http.MethodPost,
			body:       `{"operation":"add","b":5}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "operand 'a' is required",
		},
		{
			name:       "missing operand b",
			method:     http.MethodPost,
			body:       `{"operation":"add","a":10}`,
			wantStatus: http.StatusBadRequest,
			wantError:  "operand 'b' is required",
		},
		{
			name:       "method not allowed",
			method:     http.MethodGet,
			body:       "",
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  "method not allowed",
			wantAllow:  http.MethodPost,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(
				tt.method,
				"/api/v1/calculate",
				strings.NewReader(tt.body),
			)

			recorder := httptest.NewRecorder()

			handler.Calculate(recorder, request)

			// Check the response status code and other expectations
			if recorder.Code != tt.wantStatus {
				t.Errorf(
					"Calculate() status = %d, want %d",
					recorder.Code,
					tt.wantStatus,
				)
			}

			// Check the Allow header if expected
			if tt.wantAllow != "" {
				if got := recorder.Header().Get("Allow"); got != tt.wantAllow {
					t.Errorf(
						"Calculate() Allow header = %q, want %q",
						got,
						tt.wantAllow,
					)
				}
			}

			// Check the response body for result or error
			if tt.wantResult != nil {
				var response calculateResponse

				// Decode the JSON response body into the calculateResponse struct.
				if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				// Compare the result with the expected value
				if response.Result != *tt.wantResult {
					t.Errorf(
						"Calculate() result = %v, want %v",
						response.Result,
						*tt.wantResult,
					)
				}
			}

			// Check the response body for error message if expected
			if tt.wantError != "" {
				var response errorResponse

				// Decode the JSON response body into the errorResponse struct.
				if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
					t.Fatalf("failed to decode error response: %v", err)
				}

				// Compare the error message with the expected value
				if response.Error != tt.wantError {
					t.Errorf(
						"Calculate() error = %q, want %q",
						response.Error,
						tt.wantError,
					)
				}
			}
		})
	}
}

// float64Ptr is a helper function that returns a pointer to the given float64 value.
func float64Ptr(value float64) *float64 {
	return &value
}
