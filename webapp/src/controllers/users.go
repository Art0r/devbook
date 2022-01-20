package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/responses"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/user", config.ApiUrl)
	r.ParseForm()

	fmt.Println(url)
	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CatchErrorStatusCode(rw, response)
		return
	}

	responses.JSON(rw, response.StatusCode, nil)
}
