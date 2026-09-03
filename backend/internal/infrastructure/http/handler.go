package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Jcoll05/full-stack-calculator/backend/internal/domain"
)

// Calculator defines the calculation behavior required by the HTTP handler.
// Implementations receive a complete mathematical expression and return its result.
type Calculator interface {
	Calculate(expression string) (float64, error)
}

// CalculatorHandler handles HTTP requests for mathematical expression evaluation.
type CalculatorHandler struct {
	service Calculator
}

// NewCalculatorHandler creates a new CalculatorHandler with the provided calculator service.
func NewCalculatorHandler(service Calculator) *CalculatorHandler {
	return &CalculatorHandler{
		service: service,
	}
}

// calculateRequest represents the expected JSON structure for a calculation request.
type calculateRequest struct {
	Expression string `json:"expression"`
}

// calculateResponse represents the JSON structure for a successful calculation response.
type calculateResponse struct {
	Result float64 `json:"result"`
}

// errorResponse represents the JSON structure for an error response.
type errorResponse struct {
	Error string `json:"error"`
}

// Calculate handles HTTP requests for evaluating mathematical expressions.
func (h *CalculatorHandler) Calculate(w http.ResponseWriter, r *http.Request) {

	// Reject unsupported HTTP methods.
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

	// Reject requests without an expression.
	if request.Expression == "" {
		writeError(w, http.StatusBadRequest, "expression is required")
		return
	}

	// Evaluate the expression through the calculator service.
	result, err := h.service.Calculate(request.Expression)

	// Handle errors and send appropriate HTTP responses.
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrDivisionByZero),
			errors.Is(err, domain.ErrNegativeSquareRoot),
			errors.Is(err, domain.ErrInvalidExpression),
			errors.Is(err, domain.ErrMissingClosingParenthesis),
			errors.Is(err, domain.ErrUnexpectedOperator),
			errors.Is(err, domain.ErrInvalidNumber):
			writeError(w, http.StatusBadRequest, err.Error())

		default:
			writeError(w, http.StatusInternalServerError, "internal server error")
		}

		return
	}

	// Return the calculated result.
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
