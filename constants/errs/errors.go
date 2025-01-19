package errs

import "errors"

var (
	ErrEmailExists     = errors.New("email already exists")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidUser     = errors.New("invalid user data")
	ErrInvalidUUID     = errors.New("invalid UUID format")
	ErrPasswordHashing = errors.New("failed to hash password")
)
