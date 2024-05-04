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
