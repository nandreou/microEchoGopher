package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.io/nicksauth/database"
	"github.io/nicksauth/handlers/jsonrw"
	"github.io/nicksauth/logger"
	"github.io/nicksauth/models"
)

type HandlerConfig struct {
	DB     *database.DB
	Logger *logger.Logger
}

var HandlConf HandlerConfig

func NewHandlConf(logger *logger.Logger, db *database.DB) {
	HandlConf.DB = db
	HandlConf.Logger = logger
}

func (HandlConf *HandlerConfig) GetUser(w http.ResponseWriter, r *http.Request) {

	requestData, err := jsonrw.ReadJSON(r)

	//Write Request Log Start
	status, errLog := HandlConf.Logger.WriteLog(HandlConf.Logger.AuthRequestsURL, &models.AuthRequestLogModel{
		IP:    r.RemoteAddr,
		URL:   r.URL.String(),
		Email: requestData.Email,
	})
	if errLog != nil {
		log.Println(err)
	} else {
		log.Println("HTTP: /auth-request:", status)
	}
	//Write Request Log End

	if err != nil {
		log.Println(err)
		return
	}

	authenticated, err := HandlConf.DB.Auth(requestData.Email, requestData.Password)

	if err != nil {
		//Write Response Log Start
		status, errLog = HandlConf.Logger.WriteLog(HandlConf.Logger.AuthResponsesURL, &models.AuthResponseLogModel{
			IP:     r.RemoteAddr,
			URL:    r.URL.String(),
			Email:  requestData.Email,
			STATUS: 500,
			Error:  fmt.Sprintf("Error on Reading From DB: %s", err),
		})

		if errLog != nil {
			log.Println(errLog)
		} else {
			log.Println("HTTP: /auth-response:", status)
		}
		//Write Response Log End

		fmt.Println("Error", err)
		return
	}

	err = jsonrw.WriteJSON(w, authenticated, r.Header)

	if err != nil {
		//Write Response Log Start
		status, errLog = HandlConf.Logger.WriteLog(HandlConf.Logger.AuthResponsesURL, &models.AuthResponseLogModel{
			IP:     r.RemoteAddr,
			URL:    r.URL.String(),
			Email:  requestData.Email,
			STATUS: 500,
			Error:  fmt.Sprintf("Error on Reading From DB: %s", err),
		})

		if errLog != nil {
			log.Println(errLog)
		} else {
			log.Println("HTTP: /auth-response:", status)
		}
		//Write Response Log End

		log.Println("Error on Encode: ", err)
		http.Error(w, fmt.Sprintln("Server Error"), http.StatusInternalServerError)
		return

	}

	if !authenticated {
		//Write Response Log Start
		status, errLog = HandlConf.Logger.WriteLog(HandlConf.Logger.AuthResponsesURL, &models.AuthResponseLogModel{
			IP:     r.RemoteAddr,
			URL:    r.URL.String(),
			Email:  requestData.Email,
			STATUS: 401,
			Error:  fmt.Sprintf("Not Authenticated: %d", 401),
		})

		if errLog != nil {
			log.Println(errLog)
		} else {
			log.Println("HTTP: /auth-response:", status)
		}
		log.Println(r.URL.String() + " HTTP: 401")
		//Write Response Log End
		return
	}

	//Write Response Log Start
	status, errLog = HandlConf.Logger.WriteLog(HandlConf.Logger.AuthResponsesURL, &models.AuthResponseLogModel{
		IP:     r.RemoteAddr,
		URL:    r.URL.String(),
		Email:  requestData.Email,
		STATUS: 200,
		Error:  "",
	})

	if errLog != nil {
		log.Println(errLog)
	} else {
		log.Println("HTTP: /auth-response:", status)
	}
	//Write Response Log End

	log.Println(r.URL.String() + " HTTP: 200")

}
