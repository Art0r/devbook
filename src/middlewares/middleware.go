package middlewares

import (
	"devbook-api/src/authenticate"
	"devbook-api/src/responses"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(rw, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := authenticate.ValidateToken(r); err != nil {
			responses.Error(rw, http.StatusUnauthorized, err)
			return
		}
		next(rw, r)
	}
}
