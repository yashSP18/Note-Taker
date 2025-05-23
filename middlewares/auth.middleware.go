package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/yash-gkmit/NOTE-TAKER/constants"
	"github.com/yash-gkmit/NOTE-TAKER/helpers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.SendHandlerErrResponse(w, "Authorization Header Required!", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helpers.SendHandlerErrResponse(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := helpers.VerifyJWT(tokenString)
		if err != nil {
			helpers.SendHandlerErrResponse(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), constants.UserContextKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
