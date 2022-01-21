package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(rw, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if _, err := cookies.Read(r); err != nil {
			http.Redirect(rw, r, "/login", 302)
			return
		}
		next(rw, r)
	}
}
