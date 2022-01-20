package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(rw http.ResponseWriter, id, token string) error {
	data := map[string]string{
		"id":    id,
		"token": token,
	}

	encodedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
