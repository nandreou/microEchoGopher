package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.io/nicksauth/database"
	"github.io/nicksauth/handlers/jsonrw"
)

type Repository struct {
	DB *database.DB
}

var Repo Repository

func NewRepo(db *database.DB) {
	Repo.DB = db
}

func (repo *Repository) GetUser(w http.ResponseWriter, r *http.Request) {

	requestData, err := jsonrw.ReadJSON(r)

	if err != nil {
		log.Println(err)
		return
	}

	authenticated, err := Repo.DB.Auth(requestData.Email, requestData.Password)

	if err != nil {
		fmt.Println("Error", err)
	}

	err = jsonrw.WriteJSON(w, authenticated, r.Header)

	if err != nil {
		log.Println("Error on Encode: ", err)
		http.Error(w, fmt.Sprintln("Server Error"), http.StatusInternalServerError)
	}

}
