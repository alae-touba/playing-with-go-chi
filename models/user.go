package models

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
