package models

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SingUpData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
