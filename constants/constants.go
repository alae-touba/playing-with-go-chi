package constants

const (
	// Error messages
	ErrInvalidUserID                  = "Invalid user ID"
	ErrUserNotFound                   = "User not found"
	ErrInvalidRequestBody             = "Invalid request payload"
	ErrUnauthorizedInvalidCredentials = "Unauthorized. Invalid credentials"
	ErrUnauthorizedNoCredentials      = "Unauthorized. No credentials provided"
	ErrFailedToHashPassword           = "Failed to hash password"
	ErrFailedToCreateUser             = "Failed to create user"
	ErrFailedToGetUsers               = "Failed to get users"
)

const (
	// HTTP Headers
	HeaderWWWAuthenticate = "WWW-Authenticate"
	HeaderContentType     = "Content-Type"
	HeaderAuthorization   = "Authorization"
	HeaderApplicationJSON = "application/json"
)

const DefaultPort = "3005"
