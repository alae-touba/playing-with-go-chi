package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
