package middlewares

import (
	"net/http"

	"github.com/alae-touba/playing-with-go-chi/constants"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/alae-touba/playing-with-go-chi/utils"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	logger      *zap.Logger
	userService *services.UserService
}

func NewAuthMiddleware(logger *zap.Logger, userService *services.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		logger:      logger,
		userService: userService,
	}
}

func (authMiddleware *AuthMiddleware) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set(constants.HeaderWWWAuthenticate, `Basic realm="restricted"`)
			utils.RespondWithError(w, http.StatusUnauthorized, constants.ErrUnauthorizedNoCredentials)
			return
		}

		// Verify credentials
		if !authMiddleware.userService.ValidateCredentials(username, password) {
			w.Header().Set(constants.HeaderWWWAuthenticate, `Basic realm="restricted"`)
			utils.RespondWithError(w, http.StatusUnauthorized, constants.ErrUnauthorizedInvalidCredentials)
			return
		}

		authMiddleware.logger.Info("successful authentication", zap.String("username", username))
		next.ServeHTTP(w, r)
	})
}
