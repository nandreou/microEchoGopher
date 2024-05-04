package models

type JsonRequestModel struct {
	ApiKey  string `json:"apiKey"`
	Message string `json:"message"`
}

type JsonResponseModel struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type JsonLogInRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JsonLogInResponseModel struct {
	Authenticated bool   `json:"authenticated"`
	ApiKey        string `json:"apiKey"`
}
