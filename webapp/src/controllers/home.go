package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
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

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(rw, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})

}

func LoadEditPage(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	pid, err := strconv.ParseUint(params["pid"], 10, 64)
	if err != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, pid)
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

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(rw, "update-post.html", post)
}
