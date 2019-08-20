package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/kanehara/prisma-template/prisma"
)

type contextKey string

var authTokenCtxKey = contextKey("authToken")

func AuthMiddleware(client *prisma.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: authentication
			ctx := context.WithValue(r.Context(), authTokenCtxKey, "cjwwzj8b0001p0856ujul1841")
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetContextUser(ctx context.Context) (*User, error) {
	token, ok := ctx.Value(authTokenCtxKey).(string)
	if !ok {
		return nil, errors.New("Missing auth token")
	}
	u := &User{token}
	return u, nil
}
