package auth

import (
	"net/http"
)

// Middleware - header middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}
