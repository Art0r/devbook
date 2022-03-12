package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedAt time.Time `json:"createAt"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

func SearchForUser(uid uint64, r *http.Request) (User, error) {
	chUser := make(chan User)
	chFollowers := make(chan []User)
	chFollowing := make(chan []User)
	chPosts := make(chan []Post)

	go searchUserData(chUser, uid, r)
	go seachFollowers(chFollowers, uid, r)
	go seachFollowing(chFollowing, uid, r)
	go seachPosts(chPosts, uid, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case thisUser := <-chUser:
			if thisUser.Id == 0 {
				return User{}, errors.New("erro ao buscar o usuário")
			}
			user = thisUser

		case thisFollowers := <-chFollowers:
			if thisFollowers == nil {
				return User{}, errors.New("erro ao buscas seguidores")
			}
			followers = thisFollowers

		case thisFollowing := <-chFollowing:
			if thisFollowing == nil {
				return User{}, errors.New("erro ao buscar usuários seguindo")
			}
			following = thisFollowing

		case thisPosts := <-chPosts:
			if thisPosts == nil {
				return User{}, errors.New("erro ao buscar publicações")
			}
			posts = thisPosts
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func searchUserData(ch chan<- User, uid uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d", config.ApiUrl, uid)
	response, err := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		ch <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		ch <- User{}
		return
	}

	ch <- user
}

func seachFollowers(ch chan<- []User, uid uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/followers", config.ApiUrl, uid)
	response, err := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		ch <- nil
		return
	}
	defer response.Body.Close()

	var followers []User

	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		ch <- nil
		return
	}

	ch <- followers
}

func seachFollowing(ch chan<- []User, uid uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/following", config.ApiUrl, uid)
	response, err := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		ch <- nil
		return
	}
	defer response.Body.Close()

	var following []User

	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		ch <- nil
		return
	}

	ch <- following
}

func seachPosts(ch chan<- []Post, uid uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/posts", config.ApiUrl, uid)
	response, err := requests.MakeAuthRequest(r, http.MethodGet, url, nil)
	if err != nil {
		ch <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		ch <- nil
		return
	}

	ch <- posts
}
