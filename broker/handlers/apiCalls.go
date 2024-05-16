package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"guthub.io/nicksbroker/models"
)

func AuthCall(email string, password string) (*models.JsonLogInResponseModel, error) {

	body := models.JsonLogInRequestModel{Email: email, Password: password}

	byteBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, "http://192.168.1.17:8080", bytes.NewReader(byteBody))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	//responseBody, err := io.ReadAll(response.Body)

	responseBody := &models.JsonLogInResponseModel{}
	err = json.NewDecoder(response.Body).Decode(responseBody)

	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
