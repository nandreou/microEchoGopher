package jsonrw

import (
	"encoding/json"
	"net/http"

	"guthub.io/nicksbroker/models"
)

func AuthReadJSON(r *http.Request, maxBytes int64) (*models.JsonLogInRequestModel, error) {

	buffer := &models.JsonLogInRequestModel{}

	err := json.NewDecoder(r.Body).Decode(&buffer)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func ReadJSON(r *http.Request, maxBytes int64) (*models.JsonRequestModel, error) {

	buffer := &models.JsonRequestModel{}

	err := json.NewDecoder(r.Body).Decode(&buffer)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func AuthWriteJSON(w http.ResponseWriter, authData *models.JsonLogInResponseModel, headers ...http.Header) error {

	encoder := json.NewEncoder(w)

	encoder.SetIndent("", " ")

	err := encoder.Encode(authData)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func WriteJSON(w http.ResponseWriter, requestData *models.JsonRequestModel, authenticated bool, headers ...http.Header) error {

	if authenticated {
		if len(headers) > 0 {
			for key, value := range headers[0] {
				if key == "Content-Length" {
					continue
				}
				w.Header()[key] = value
			}
		}

		resp := &models.JsonResponseModel{
			Message: "User Authenticated: Message Succefully Received",
			Data:    requestData.Message,
		}

		encoder := json.NewEncoder(w)

		encoder.SetIndent("", " ")

		err := encoder.Encode(resp)

		if err != nil {
			http.Error(w, "Error: ", http.StatusInternalServerError)
			return err
		} else {
			return nil
		}
	} else {
		http.Error(w, `Not Authenticated: 401 GO TO "/login" to get an ApiKey`, 401)
		return nil
	}
}
