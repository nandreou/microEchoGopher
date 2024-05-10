package models

import "time"

// REQUEST - RESPONSE MODEL
type JsonRequestModel struct {
	ApiKey  string `json:"apiKey"`
	Email   string `json:"email"`
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

// LOGER MODELS
type BrokerRequestLogModel struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	IP        string    `bson:"ip" json:"ip"`
	URL       string    `bson:"url" json:"url"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type BrokerResponseLogModel struct {
	ID     string `bson:"_id,omitempty" json:"id,omitempty"`
	IP     string `bson:"ip" json:"ip"`
	STATUS int    `bson:"status" json:"status"`
	URL    string `bson:"url" json:"url"`
	Error  string `bson:"error" json:"error"`
	Email  string `bson:"email" json:"email"`
	// CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
