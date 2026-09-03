package http

import "net/http"

// corsMiddleware is a middleware function that adds CORS headers to the HTTP response.
func corsMiddleware(next http.Handler) http.Handler {
	// Return a new handler that wraps the next handler and adds CORS headers.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests by responding with a 204 No Content status.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}