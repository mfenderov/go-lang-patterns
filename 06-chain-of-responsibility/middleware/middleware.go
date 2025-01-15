package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Logging request")
		next.ServeHTTP(w, r)
		log.Printf("Request completed in %v", time.Since(start))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authenticating user")
		next.ServeHTTP(w, r)
	})
}

func CompressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Println("Response compressed")
	})
}
