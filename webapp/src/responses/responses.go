package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	if err := json.NewEncoder(rw).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func CatchErrorStatusCode(rw http.ResponseWriter, r *http.Response) {
	var err Error
	json.NewDecoder(r.Body).Decode(&err)
	JSON(rw, r.StatusCode, err)
}
