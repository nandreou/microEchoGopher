package handlers

import (
	"log"
	"net/http"

	"github.io/nickslogging/handlers/validators"
	"github.io/nickslogging/mongodb"
)

type Conf struct {
	mongo *mongodb.DB
}

var HandlersConf Conf

func NewHanldersConf(dbs *mongodb.DB) {
	HandlersConf.mongo = dbs
}

func (handlerConf *Conf) InsertBrokerReqLog(w http.ResponseWriter, r *http.Request) {

	brokerRequest, err := validators.ValidateBrokerRequestBody(w, r)

	if err != nil {
		http.Error(w, "Server Error", 500)
		log.Println(err)
		return
	}

	err = handlerConf.mongo.InsertToDB(brokerRequest, handlerConf.mongo.Mongo.BrokerRequestsCollection)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(r.URL.String() + " HTTP: 200")
}

func (handlerConf *Conf) InsertBrokerRespLog(w http.ResponseWriter, r *http.Request) {
	brokerResponse, err := validators.ValidateBrokerResponseBody(w, r)

	if err != nil {
		http.Error(w, "Server Error", 500)
		log.Println(err)
		return
	}

	err = handlerConf.mongo.InsertToDB(brokerResponse, handlerConf.mongo.Mongo.BrokerResponsesCollection)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(r.URL.String() + " HTTP: 200")
}

func (handlerConf *Conf) InsertAuthReqLog(w http.ResponseWriter, r *http.Request) {
	authRequest, err := validators.ValidateAuthRequestBody(w, r)

	if err != nil {
		http.Error(w, "Server Error", 500)
		log.Println(err)
		return
	}

	err = handlerConf.mongo.InsertToDB(authRequest, handlerConf.mongo.Mongo.AuthRequestsCollection)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(r.URL.String() + " HTTP: 200")
}

func (handlerConf *Conf) InsertAuthRespLog(w http.ResponseWriter, r *http.Request) {
	authRequest, err := validators.ValidateAuthResponseBody(w, r)

	if err != nil {
		http.Error(w, "Server Error", 500)
		log.Println(err)
		return
	}

	err = handlerConf.mongo.InsertToDB(authRequest, handlerConf.mongo.Mongo.AuthResponsesCollection)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(r.URL.String() + " HTTP: 200")
}
