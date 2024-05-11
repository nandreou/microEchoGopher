package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	CreatedAt any    //time.Time `json:"created_at"`
	UpdatedAt any    //time.Time `json:"updated_at"`
}

type JsonResponseModel struct {
	Auth   bool   `json:"authenticated"`
	ApiKey string `json:"apiKey"`
}

type JsonRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LOGER MODELS
type AuthRequestLogModel struct {
	ID    string `bson:"_id,omitempty" json:"id,omitempty"`
	IP    string `bson:"ip" json:"ip"`
	URL   string `bson:"url" json:"url"`
	Email string `bson:"email" json:"email"`
}

type AuthResponseLogModel struct {
	ID     string `bson:"_id,omitempty" json:"id,omitempty"`
	IP     string `bson:"ip" json:"ip"`
	STATUS int    `bson:"status" json:"status"`
	URL    string `bson:"url" json:"url"`
	Error  string `bson:"error" json:"error"`
	Email  string `bson:"email" json:"email"`
}
