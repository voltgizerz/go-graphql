package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-graphql/logger"
)

type contextKey string

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
const (
	CONTEXT_USER contextKey = "CTX_USER"
)

// User - a stand-in for our database backed user object
type User struct {
	UserID  int
	Name    string
	IsAdmin bool
}

// Middleware - handle auth data
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := strconv.Atoi(r.Header.Get("user_id"))
			if err != nil {
				logger.Log.Error("Header user_id is invalid or not set!")
			}

			user := &User{UserID: userID, Name: "felix", IsAdmin: true}

			// put it in context
			ctx := context.WithValue(r.Context(), CONTEXT_USER, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(CONTEXT_USER).(*User)
	return raw
}
