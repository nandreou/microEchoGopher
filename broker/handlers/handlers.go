package handlers

import (
	"fmt"
	"log"
	"net/http"

	"guthub.io/nicksbroker/config"
	"guthub.io/nicksbroker/database"
	"guthub.io/nicksbroker/handlers/jsonrw"
)

type Repository struct {
	App *config.App
	DB  *database.DB
}

var Repo Repository

func NewRepo(a *config.App, db *database.DB) {
	Repo.App = a
	Repo.DB = db
}

func (repo *Repository) LogIn(w http.ResponseWriter, r *http.Request) {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	requestData, err := jsonrw.AuthReadJSON(r, int64(maxBytes))

	responseData, err := AuthCall(requestData.Email, requestData.Password)

	if err != nil {
		http.Error(w, "Sorry Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if !responseData.Authenticated {
		if err := jsonrw.AuthWriteJSON(w, responseData, r.Header); err != nil {
			http.Error(w, "Sorry Error", http.StatusInternalServerError)
			log.Println("Error on Writing json: ", err)
		}
		return
	} else {
		_, err = repo.DB.WriteApiKeyToDB(requestData.Email, responseData.ApiKey)

		if err != nil {
			http.Error(w, "Sorry Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if err := jsonrw.AuthWriteJSON(w, responseData, r.Header); err != nil {
			http.Error(w, "Sorry Error", http.StatusInternalServerError)
			log.Println("Error on Writing json: ", err)
		}
	}
}

func (repo *Repository) EchoHandler(w http.ResponseWriter, r *http.Request) {

	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	requestData, err := jsonrw.ReadJSON(r, int64(maxBytes))

	if err != nil {
		log.Println("Error on Reading Data: ", err)
		http.Error(w, `Bad Request`, http.StatusBadRequest)
		return
	}

	authenticated, err := repo.DB.ApiKeyValidate(requestData.ApiKey)

	if err != nil {
		if fmt.Sprintf("%s", err) == "sql: no rows in result set" {
			http.Error(w, `Not Authenticated: 401 GO TO "/login" to get an ApiKey`, 401)
			return
		}

		http.Error(w, "Sorry Error", http.StatusInternalServerError)
		log.Println("Error on Fetching From Database ApiKeyValidate(): ", err)
		return
	}

	if err != jsonrw.WriteJSON(w, requestData, authenticated, r.Header) {
		log.Println("Error on Writing json: ", err)
	}
}
