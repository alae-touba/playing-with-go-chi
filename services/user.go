package services

import (
	"fmt"
	"github.com/alae-touba/playing-with-go-chi/models"
	"github.com/alae-touba/playing-with-go-chi/security"
)

type UserService struct {
	users []models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: []models.User{},
	}
}

func (s *UserService) GetUsers(name string) []models.UserResponse {
	var filtered []models.UserResponse
	for _, user := range s.users {
		if name == "" || user.Name == name {
			filtered = append(filtered, models.UserResponse{
				ID:   user.ID,
				Name: user.Name,
			})
		}
	}
	return filtered
}

func (s *UserService) GetUser(id string) *models.UserResponse {
	for _, user := range s.users {
		if user.ID == id {
			return &models.UserResponse{ID: user.ID, Name: user.Name}
		}
	}
	return nil
}

func (s *UserService) ValidateCredentials(username, password string) bool {
	user := s.findUserByUsername(username)
	if user == nil {
		return false
	}

	return security.VerifyPassword(password, user.Password)
}

func (s *UserService) CreateUser(name, password string) (*models.UserResponse, error) {
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		ID:       fmt.Sprintf("%d", len(s.users)+1),
		Name:     name,
		Password: hashedPassword,
	}

	s.users = append(s.users, newUser)
	return &models.UserResponse{
		ID:   newUser.ID,
		Name: newUser.Name,
	}, nil
}

func (s *UserService) findUserByUsername(userame string) *models.User {
	for _, user := range s.users {
		if user.Name == userame {
			return &user
		}
	}
	return nil
}
