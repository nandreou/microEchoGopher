package models

import "time"

type BrokerRequestModel struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	IP        string    `bson:"ip" json:"ip"`
	URL       string    `bson:"url" json:"url"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type BrokerResponseModel struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	IP        string    `bson:"ip" json:"ip"`
	STATUS    int       `bson:"status" json:"status"`
	URL       string    `bson:"url" json:"url"`
	Error     string    `bson:"error" json:"error"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type AuthRequestModel struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	IP        string    `bson:"ip" json:"ip"`
	URL       string    `bson:"url" json:"url"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type AuthResponseModel struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	IP        string    `bson:"ip" json:"ip"`
	STATUS    int       `bson:"status" json:"status"`
	URL       string    `bson:"url" json:"url"`
	Error     string    `bson:"error" json:"error"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
