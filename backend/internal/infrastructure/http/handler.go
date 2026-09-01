package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Jcoll05/full-stack-calculator/backend/internal/domain"
)

// Calculator defines the calculation behavior required by the HTTP handler.
type Calculator interface {
	Calculate(operation domain.Operation, a, b float64) (float64, error)
}

// CalculatorHandler handles HTTP requests for calculator operations.
type CalculatorHandler struct {
	service Calculator
}

// NewCalculatorHandler creates a new instance of CalculatorHandler with the provided CalculatorService.
func NewCalculatorHandler(service Calculator) *CalculatorHandler {
	return &CalculatorHandler{
		service: service,
	}
}

// calculateRequest represents the expected JSON structure for a calculation request.
type calculateRequest struct {
	Operation string   `json:"operation"`
	A         *float64 `json:"a"` // Use pointer (*) to distinguish between zero value and missing field.
	B         *float64 `json:"b"`
}

// calculateResponse represents the JSON structure for a successful calculation response.
type calculateResponse struct {
	Result float64 `json:"result"`
}

// errorResponse represents the JSON structure for an error response.
type errorResponse struct {
	Error string `json:"error"`
}

// Calculate handles the HTTP request for performing a calculation.
func (h *CalculatorHandler) Calculate(w http.ResponseWriter, r *http.Request) {

	// Ensure the request method is POST; otherwise, return a 405 Method Not Allowed error.
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var request calculateRequest

	// Decode the JSON request body into the calculateRequest struct.
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Validate that an operation was provided.
	if request.Operation == "" {
		writeError(w, http.StatusBadRequest, "operation is required")
		return
	}

	// Validate that operand 'a' was provided.
	if request.A == nil {
		writeError(w, http.StatusBadRequest, "operand 'a' is required")
		return
	}

	operation := domain.Operation(request.Operation)

	// Binary operations require operand 'b'.
	switch operation {
	case domain.Add,
		domain.Subtract,
		domain.Multiply,
		domain.Divide,
		domain.Power:
		if request.B == nil {
			writeError(w, http.StatusBadRequest, "operand 'b' is required")
			return
		}
	}

	// Unary operations do not require operand 'b'.
	var b float64
	if request.B != nil {
		b = *request.B
	}

	// Perform the calculation using the CalculatorService.
	result, err := h.service.Calculate(
		operation,
		*request.A,
		b,
	)

	// Handle errors and send appropriate HTTP responses.
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUnsupportedOperation),
			errors.Is(err, domain.ErrDivisionByZero),
			errors.Is(err, domain.ErrNegativeSquareRoot):
			writeError(w, http.StatusBadRequest, err.Error())

		default:
			writeError(w, http.StatusInternalServerError, "internal server error")
		}

		return
	}

	// Send a successful response with the calculation result.
	writeJSON(w, http.StatusOK, calculateResponse{
		Result: result,
	})
}

// writeJSON sends a JSON response with the specified status code and data.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

// writeError sends a JSON error response with the specified status code and error message.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{
		Error: message,
	})
}
