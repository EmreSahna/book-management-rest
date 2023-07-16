package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request from: %s to %s\n", r.RemoteAddr, r.URL)
		next.ServeHTTP(w, r)
	})
}
