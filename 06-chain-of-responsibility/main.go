package main

import (
	"chain-of-responsibillity/middleware"
	"log"
	"net/http"
)

// Middleware type
type Middleware func(http.Handler) http.Handler

// Request Handler
func requestHandler(_ http.ResponseWriter, _ *http.Request) {
	log.Println("Hello, World! Your request has been processed.")
}

// Chaining Middleware
func addMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- { // Apply in reverse order
		handler = middlewares[i](handler)
	}
	return handler
}

func main() {
	// Define the requestHandler handler
	requestHandler := http.HandlerFunc(requestHandler)

	// Chain middlewares
	handler := addMiddleware(
		requestHandler,
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		middleware.CompressionMiddleware,
	)

	// Simulation of a request
	handler.ServeHTTP(nil, nil)
}
