package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func LoadLoginScreen(rw http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(rw, "login.html", nil)
}

func LoadUserSigninPage(rw http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(rw, "signup.html", nil)
}
