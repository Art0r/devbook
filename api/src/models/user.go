package models

import (
	"devbook-api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}
	return nil
}
func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("name field is empty")
	}

	if user.Email == "" {
		return errors.New("email field is empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("inserted email is invalid")
	}

	if user.Nick == "" {
		return errors.New("nick field is empty")
	}

	if stage == "registration" && user.Password == "" {
		return errors.New("Password field is empty")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)

	if stage == "registration" {
		passwordHash, err := security.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}
	return nil
}
