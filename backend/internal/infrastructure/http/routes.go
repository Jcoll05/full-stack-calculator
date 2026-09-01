package http

import "net/http"

// NewRouter creates and configures the HTTP router for the calculator API.
func NewRouter(handler *CalculatorHandler) http.Handler {
	// Create a new ServeMux to handle incoming HTTP requests.
	mux := http.NewServeMux()

	// Register the Calculate handler for the /api/v1/calculate endpoint.
	mux.HandleFunc("/api/v1/calculate", handler.Calculate)

	return mux // Return the configured router to be used by the HTTP server.
}
