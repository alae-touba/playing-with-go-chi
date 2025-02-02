package errs

import "errors"

var (
	ErrEmailExists     = errors.New("email already exists")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidUser     = errors.New("invalid user data")
	ErrPasswordHashing = errors.New("failed to hash password")
	ErrValidationFailed = errors.New("database validation failed")
	
	ErrTopicNotFound = errors.New("topic not found")

	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidUUID     = errors.New("invalid UUID format")
)
