package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/user", config.ApiUrl)
	r.ParseForm()

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

func UnfollowUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/user/%d/unfollow", config.ApiUrl, uid)
	response, err := requests.MakeAuthRequest(r, http.MethodPost, url, nil)
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

func FollowUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/user/%d/follow", config.ApiUrl, uid)
	response, err := requests.MakeAuthRequest(r, http.MethodPost, url, nil)
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

func EditUser(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if err != nil {
		fmt.Println(err)
		responses.JSON(rw, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	uid, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/user/%d", config.ApiUrl, uid)

	response, err := requests.MakeAuthRequest(r, http.MethodPut, url, bytes.NewBuffer(user))
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

func UpdatePassword(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	password, err := json.Marshal(map[string]string{
		"this": r.FormValue("oldPassword"),
		"new":  r.FormValue("newPassword"),
	})
	if err != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	uid, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/user/%d/changepassword", config.ApiUrl, uid)
	res, err := requests.MakeAuthRequest(r, http.MethodPost, url, bytes.NewBuffer(password))
	if err != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.CatchErrorStatusCode(rw, res)
		return
	}

	responses.JSON(rw, res.StatusCode, nil)
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	uid, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/user/%d", config.ApiUrl, uid)

	res, err := requests.MakeAuthRequest(r, http.MethodDelete, url, nil)
	if err != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.CatchErrorStatusCode(rw, res)
		return
	}

	responses.JSON(rw, res.StatusCode, nil)
}
