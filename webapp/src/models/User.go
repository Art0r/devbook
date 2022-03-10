package models

import "time"

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
