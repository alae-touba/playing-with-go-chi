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

	// time.Time has always a value and cannot be nil
	// the value can be:
	//   - zero value (January 1, year 1, 00:00:00 UTC)
	//   - valid time (any other timestamp)
	// if we do not use pointer we will always have a value wether it is zero or valid time
	// Using *time.Time (pointer) allows for:
	//   - nil (field will be omitted from JSON)
	//   - valid time reference (will be included in JSON)
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ImageName string `json:"image_name,omitempty"`
}
