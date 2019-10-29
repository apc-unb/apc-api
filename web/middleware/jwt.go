package middleware

import (
	"github.com/apc-unb/apc-api/auth"
	"github.com/apc-unb/apc-api/web/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func SetMiddlewareJSON() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}

func SetMiddlewareAuthentication() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var hostname string
			var err error

			if err = auth.CheckToken(r); err != nil {
				logrus.Infof(err.Error())
				utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			if hostname, err = os.Hostname(); err != nil {
				logrus.Infof(err.Error())
				utils.RespondWithError(w, http.StatusInternalServerError, "Internal error")
				return
			}

			w.Header().Set("X-ContainerId", hostname)
			next.ServeHTTP(w, r)
		})
	}
}