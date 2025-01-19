package models

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	ImageName string    `json:"image_name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	//normally zero values are omitted from JSON if we use omitempty
	//but for some reason, time.Time zero value is showed even if we use omitempty
	//to avoid this we use pointer (nil field will be omitted from JSON)
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ImageName string `json:"image_name"`
}
