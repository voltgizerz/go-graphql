package auth

import (
	"net/http"
)

// HeaderMiddleware - add header as context
func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
