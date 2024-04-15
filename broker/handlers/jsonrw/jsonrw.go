package jsonrw

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResponseModel struct {
	Err     bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ReadJSON(r *http.Request, maxBytes int64) (any, error) {

	var buffer any

	err := json.NewDecoder(r.Body).Decode(&buffer)

	if err != nil {
		fmt.Println("DECODE ", err)
		return "", err
	}

	return buffer, nil
}

func WriteJSON(w http.ResponseWriter, data any, headers ...http.Header) error {

	resp := jsonResponseModel{
		false,
		"Message Succefully Received",
		data,
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
		fmt.Println("Error on Encode: ", err)

		resp = jsonResponseModel{
			true,
			"Message Failed to be Received",
			fmt.Sprintln(err),
		}
		http.Error(w, fmt.Sprintln("{\n"+"	Error:", resp.Err, "\n	Message:", resp.Message, "\n	Data: ", resp.Data, "\n}"), http.StatusInternalServerError)
		return err
	} else {
		return nil
	}

}
