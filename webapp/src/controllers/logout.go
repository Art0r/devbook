package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

func Logout(rw http.ResponseWriter, r *http.Request) {
	cookies.Delete(rw)
	http.Redirect(rw, r, "/login", http.StatusFound)
}
