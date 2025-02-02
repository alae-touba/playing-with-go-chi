package models

import "time"

// for create
type TopicRequest struct {
    Name        string `json:"name" validate:"required"`
    Description string `json:"description"`
    ImageName   string `json:"image_name"`
    UserID      string `json:"user_id" validate:"required,uuid"`
}

// for update
type TopicUpdateRequest struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    ImageName   string `json:"image_name"`
    UserID      string `json:"user_id"`
}

type TopicResponse struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    ImageName   string    `json:"image_name,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    UserID      string    `json:"user_id"`
}