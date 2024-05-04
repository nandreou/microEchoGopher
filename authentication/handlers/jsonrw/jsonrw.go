package jsonrw

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.io/nicksauth/models"
)

func ReadJSON(r *http.Request) (*models.JsonRequestModel, error) {

	buffer := &models.JsonRequestModel{}

	err := json.NewDecoder(r.Body).Decode(buffer)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func WriteJSON(w http.ResponseWriter, authenticated bool, headers ...http.Header) error {

	resp := &models.JsonResponseModel{}

	if authenticated {

		apiKey := uuid.New()

		*resp = models.JsonResponseModel{
			Auth:   authenticated,
			ApiKey: apiKey.String(),
		}

	} else {
		*resp = models.JsonResponseModel{
			Auth:   authenticated,
			ApiKey: "nil",
		}
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			if key == "Content-Length" {
				continue
			}
			w.Header()[key] = value
		}
	}

	encoder := json.NewEncoder(w)

	encoder.SetIndent("", " ")

	err := encoder.Encode(resp)

	if err != nil {
		return err
	} else {
		return nil
	}

}
