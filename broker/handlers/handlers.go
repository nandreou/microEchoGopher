package handlers

import (
	"fmt"
	"net/http"

	"guthub.io/nicksbroker/config"
	"guthub.io/nicksbroker/handlers/jsonrw"
)

type Repository struct {
	App *config.App
}

var Repo Repository

func NewRepo(a *config.App) {
	Repo.App = a
}

func (repo *Repository) EchoHandler(w http.ResponseWriter, r *http.Request) {

	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	data, err := jsonrw.ReadJSON(r, int64(maxBytes))

	if err != nil {
		fmt.Println("Error on Reading Data: ", err)
		w.Write([]byte(fmt.Sprintln(err)))
		return
	}

	if err != jsonrw.WriteJSON(w, data, r.Header) {
		fmt.Println("Error on Writing json: ", err)
	}
}
