package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

func Signin(rw http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/login", config.ApiUrl)
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	response, err := http.Post(
		url, "application/json", bytes.NewBuffer(user),
	)

	if err != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CatchErrorStatusCode(rw, response)
		return
	}

	var authData models.AuthData
	if err = json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	if err = cookies.Save(rw, authData.Id, authData.Token); err != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	responses.JSON(rw, http.StatusOK, nil)

}
