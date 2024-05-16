package logger

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Logger struct {
	AuthRequestsURL  string
	AuthResponsesURL string
}

func SetUpLogger() *Logger {
	return &Logger{
		"http://192.168.1.17:8081/auth-request",
		"http://192.168.1.17:8081/auth-response",
	}
}

func (logger *Logger) WriteLog(url string, body any) (int, error) {
	byteBody, err := json.Marshal(body)

	if err != nil {
		return 500, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(byteBody))

	if err != nil {
		return 500, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Do(req)

	if err != nil {
		return response.StatusCode, err
	}

	return response.StatusCode, nil
}
