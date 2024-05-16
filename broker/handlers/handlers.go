package handlers

import (
	"fmt"
	"log"
	"net/http"

	"guthub.io/nicksbroker/database"
	"guthub.io/nicksbroker/handlers/jsonrw"
	"guthub.io/nicksbroker/logger"
	"guthub.io/nicksbroker/models"
)

type HandlerConfig struct {
	Logger *logger.Logger
	DB     *database.DB
}

var HandlConf HandlerConfig

func NewHandlConf(a *logger.Logger, db *database.DB) {
	HandlConf.Logger = a
	HandlConf.DB = db
}

func (HandlConf *HandlerConfig) LogIn(w http.ResponseWriter, r *http.Request) {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	requestData, err := jsonrw.AuthReadJSON(r, int64(maxBytes))

	if err != nil {
		http.Error(w, "Sorry Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	responseData, err := AuthCall(requestData.Email, requestData.Password)

	if err != nil {
		http.Error(w, "Sorry Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if !responseData.Authenticated {
		w.WriteHeader(http.StatusUnauthorized)

		if err := jsonrw.AuthWriteJSON(w, responseData, r.Header); err != nil {
			log.Println("Error on Writing json: ", err)
		}
		return
	} else {
		_, err = HandlConf.DB.WriteApiKeyToDB(requestData.Email, responseData.ApiKey)

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

func (HandlConf *HandlerConfig) EchoHandler(w http.ResponseWriter, r *http.Request) {

	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	requestData, err := jsonrw.ReadJSON(r, int64(maxBytes))

	//Write Request Log Start
	status, errLog := HandlConf.Logger.WriteLog(HandlConf.Logger.BrokerRequestsURL, &models.BrokerRequestLogModel{
		IP:  r.RemoteAddr,
		URL: r.URL.String(),
	})
	if errLog != nil {
		log.Println(err)
	} else {
		log.Println("HTTP: /broker-request:", status)
	}
	//Write Request Log End

	if err != nil {
		//Write Response Log Start
		status, errLog = HandlConf.Logger.WriteLog(HandlConf.Logger.BrokerResponsesURL, &models.BrokerResponseLogModel{
			IP:     r.RemoteAddr,
			URL:    r.URL.String(),
			STATUS: 500,
			Error:  fmt.Sprintf("Error on ReadingJSON: %s", err),
		})

		if errLog != nil {
			log.Println(errLog)
		} else {
			log.Println("HTTP: /broker-response:", status)
		}
		//Write Response Log End

		log.Println("Error on Reading Data: ", err)
		http.Error(w, `Error`, http.StatusInternalServerError)
		return
	}

	authenticated, err := HandlConf.DB.ApiKeyValidate(requestData.ApiKey)

	if err != nil {
		if fmt.Sprintf("%s", err) == "sql: no rows in result set" {
			http.Error(w, `Not Authenticated: 401 GO TO "/login" to get an ApiKey`, 401)

			//Write Response Log Start
			status, errLog = HandlConf.Logger.WriteLog(HandlConf.Logger.BrokerResponsesURL, &models.BrokerResponseLogModel{
				IP:     r.RemoteAddr,
				URL:    r.URL.String(),
				STATUS: 401,
				Error:  `Not Authenticated: 401 GO TO "/login" to get an ApiKey`,
			})

			if errLog != nil {
				log.Println(errLog)
			} else {
				log.Println("HTTP: /broker-response:", status)
			}
			//Write Response Log End

			log.Println("Not Authenticated: 401")
			return
		}

		//Write Response Log Start
		status, errLog = HandlConf.Logger.WriteLog(HandlConf.Logger.BrokerResponsesURL, &models.BrokerResponseLogModel{
			IP:     r.RemoteAddr,
			URL:    r.URL.String(),
			STATUS: 500,
			Error:  fmt.Sprintf("Error on Fetching From Database ApiKeyValidate(): %s", err),
		})

		if errLog != nil {
			log.Println(errLog)
		} else {
			log.Println("HTTP: /broker-response:", status)
		}
		//Write Response Log End

		http.Error(w, "Sorry Error", http.StatusInternalServerError)
		log.Println("Error on Fetching From Database ApiKeyValidate(): ", err)
		return
	}

	if err != jsonrw.WriteJSON(w, requestData, authenticated, r.Header) {

		//Write Response Log Start
		status, errLog := HandlConf.Logger.WriteLog(HandlConf.Logger.BrokerResponsesURL, &models.BrokerResponseLogModel{
			IP:     r.RemoteAddr,
			URL:    r.URL.String(),
			STATUS: 500,
			Error:  fmt.Sprintf("Error on Writing json: %s", err),
		})

		if errLog != nil {
			log.Println(errLog)
		} else {
			log.Println("HTTP: /broker-response:", status)
		}
		//Write Response Log End

		log.Println("Error on Writing json: ", err)
		return
	}

	//Write Response Log Start
	status, errLog = HandlConf.Logger.WriteLog(HandlConf.Logger.BrokerResponsesURL, &models.BrokerResponseLogModel{
		IP:     r.RemoteAddr,
		URL:    r.URL.String(),
		STATUS: 200,
		Error:  "",
	})

	if errLog != nil {
		log.Println(errLog)
	} else {
		log.Println("HTTP: /broker-response:", status)
	}
	//Write Response Log End

	log.Println(r.URL.String() + " HTTP: 200")
}
