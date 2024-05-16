package validators

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.io/nickslogging/models"
)

func ValidateBrokerRequestBody(w http.ResponseWriter, r *http.Request) (*models.BrokerRequestModel, error) {
	brokerRequest := &models.BrokerRequestModel{}

	err := json.NewDecoder(r.Body).Decode(brokerRequest)

	if err != nil {
		return nil, err
	}

	if brokerRequest.IP == "" || brokerRequest.URL == "" || brokerRequest.Email == "" {
		http.Error(w, "Request Data Not Valid", 400)
		return nil, errors.New("Data From Body Not Valid")
	}

	brokerRequest.CreatedAt = time.Now()

	return brokerRequest, nil
}

func ValidateBrokerResponseBody(w http.ResponseWriter, r *http.Request) (*models.BrokerResponseModel, error) {
	brokerResponse := &models.BrokerResponseModel{}

	err := json.NewDecoder(r.Body).Decode(brokerResponse)

	if err != nil {
		return nil, err
	}

	if brokerResponse.IP == "" || brokerResponse.URL == "" || brokerResponse.Email == "" || brokerResponse.STATUS <= 0 {
		http.Error(w, "Request Data Not Valid", 400)
		return nil, errors.New("Data From Body Not Valid")
	}

	brokerResponse.CreatedAt = time.Now()

	return brokerResponse, nil
}

func ValidateAuthRequestBody(w http.ResponseWriter, r *http.Request) (*models.AuthRequestModel, error) {
	authRequest := &models.AuthRequestModel{}

	err := json.NewDecoder(r.Body).Decode(authRequest)

	if err != nil {
		return nil, err
	}

	if authRequest.IP == "" || authRequest.URL == "" || authRequest.Email == "" {
		http.Error(w, "Request Data Not Valid", 400)
		return nil, errors.New("Data From Body Not Valid")
	}

	authRequest.CreatedAt = time.Now()

	return authRequest, nil
}

func ValidateAuthResponseBody(w http.ResponseWriter, r *http.Request) (*models.AuthResponseModel, error) {
	authResponse := &models.AuthResponseModel{}

	err := json.NewDecoder(r.Body).Decode(authResponse)

	if err != nil {
		return nil, err
	}

	if authResponse.IP == "" || authResponse.URL == "" || authResponse.Email == "" || authResponse.STATUS <= 0 {
		http.Error(w, "Request Data Not Valid", 400)
		return nil, errors.New("Data From Body Not Valid")
	}

	authResponse.CreatedAt = time.Now()

	return authResponse, nil
}
