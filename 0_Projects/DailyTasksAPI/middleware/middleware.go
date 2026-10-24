package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger middleware function
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the actual handler
		next(w, r)

		// Log method, path, and how long it took
		log.Printf("[%s] %s | Took: %v", r.Method, r.URL.Path, time.Since(start))
	}
}
