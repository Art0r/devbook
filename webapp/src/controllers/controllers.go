package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

func LoadLoginScreen(rw http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(rw, r, "/home", http.StatusFound)
		return
	}
	utils.ExecuteTemplate(rw, "login.html", nil)
}

func LoadUserSigninPage(rw http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(rw, "signup.html", nil)
}

func LoadUserPage(rw http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/user?user=%s", config.ApiUrl, nameOrNick)

	response, err := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CatchErrorStatusCode(rw, response)
		return
	}

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(rw, "users.html", users)
}
