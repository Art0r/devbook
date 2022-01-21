package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

func LoadHomePage(rw http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.ApiUrl)

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

	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(rw, "home.html", posts)
}
