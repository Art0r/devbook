package controllers

import (
	"devbook-api/src/authenticate"
	"devbook-api/src/db"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(rw http.ResponseWriter, r *http.Request) {
	uid, err := authenticate.ExtractUserID(r)
	if err != nil {
		responses.Error(rw, http.StatusUnauthorized, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err := json.Unmarshal(body, &post); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	post.AuthorId = uid

	if err = post.Prepare(); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromPosts(db)

	post.Id, err = repository.Create(post)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusCreated, post)
}

func SearchPosts(rw http.ResponseWriter, r *http.Request) {
	uid, err := authenticate.ExtractUserID(r)
	if err != nil {
		responses.Error(rw, http.StatusUnauthorized, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromPosts(db)
	posts, err := repository.Search(uid)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, posts)
}

func SearchPost(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromPosts(db)

	post, err := repository.SearchById(pid)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, post)
}

func UpdatePost(rw http.ResponseWriter, r *http.Request) {
	uid, err := authenticate.ExtractUserID(r)
	if err != nil {
		responses.Error(rw, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	pid, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromPosts(db)
	old, err := repository.SearchById(pid)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	if old.AuthorId != uid {
		responses.Error(rw, http.StatusForbidden, errors.New("a publicação pertence a outro usuário"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err := json.Unmarshal(body, &post); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(pid, post); err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusNoContent, nil)
}

func DeletePost(rw http.ResponseWriter, r *http.Request) {
	uid, err := authenticate.ExtractUserID(r)
	if err != nil {
		responses.Error(rw, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	pid, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromPosts(db)
	old, err := repository.SearchById(pid)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	if old.AuthorId != uid {
		responses.Error(rw, http.StatusForbidden, errors.New("a publicação pertence a outro usuário"))
		return
	}

	if err = repository.Delete(pid); err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusNoContent, nil)
}

func SearchByUsers(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromPosts(db)
	posts, err := repository.SearchByUser(uid)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, posts)
}

func Like(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromPosts(db)
	if err = repository.Like(pid); err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusNoContent, nil)
}

func Unlike(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryFromPosts(db)
	if err = repository.Unlike(pid); err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusNoContent, nil)
}
