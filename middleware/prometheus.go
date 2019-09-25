package middleware

import (
	"net/http"

	"github.com/VerasThiago/api/metrics"
	"github.com/gorilla/mux"
)

func GetPrometheusMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.String() != "/metrics" {
				metrics.InvocationCount.Inc()
			}
			next.ServeHTTP(w, r)
		})
	}
}
