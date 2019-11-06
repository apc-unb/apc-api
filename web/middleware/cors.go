package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetCorsMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := []string{"http://react.localhost", "http://localhost:5000", "http://localhost:4000"}
			origin := r.Header.Get("Origin")
			for _, value := range allowedOrigins {
				if value == origin {
					w.Header().Set("Access-Control-Allow-Origin", origin)
				}
			}
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Expose-Headers", "X-Container-Id")
			if r.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
