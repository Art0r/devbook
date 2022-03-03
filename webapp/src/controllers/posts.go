package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.ApiUrl)
	response, err := requests.MakeAuthRequest(r, http.MethodPost, url, bytes.NewBuffer(post))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CatchErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid, err := strconv.ParseUint(params["pid"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.ApiUrl, pid)
	response, err := requests.MakeAuthRequest(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CatchErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid, err := strconv.ParseUint(params["pid"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/dislike", config.ApiUrl, pid)
	response, err := requests.MakeAuthRequest(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CatchErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid, err := strconv.ParseUint(params["pid"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	r.ParseForm()
	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, pid)
	response, err := requests.MakeAuthRequest(r, http.MethodPut, url, bytes.NewBuffer(post))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.CatchErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
