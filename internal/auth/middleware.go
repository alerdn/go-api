package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/alerdn/go-api/internal/shared"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token ausente ou mal formatado", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ValidarToken(tokenStr)
		if err != nil {
			http.Error(w, "Token inválido: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), shared.UsuarioIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
