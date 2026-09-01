package main

import (
	"log"
	"net/http"

	"github.com/Jcoll05/full-stack-calculator/backend/internal/application"
	calculatorHTTP "github.com/Jcoll05/full-stack-calculator/backend/internal/infrastructure/http"
)

// main is the entry point of the calculator backend application.
func main() {
	// Initialize the application service.
	service := application.NewCalculatorService()

	// Initialize the HTTP handler with the application service.
	handler := calculatorHTTP.NewCalculatorHandler(service)

	// Initialize the HTTP router with the handler.
	router := calculatorHTTP.NewRouter(handler)

	// Start the HTTP server.
	const address = ":8080"

	log.Printf("Calculator backend listening on %s", address)

	// Start the HTTP server and log any errors that occur.
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}