package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Smith"},
}

const (
	ErrInvalidUserID                  = "Invalid user ID"
	ErrUserNotFound                   = "User not found"
	ErrInvalidRequestBody             = "Invalid request payload"
	ErrUnauthorizedInvalidCredentials = "Unauthorized. Invalid credentials"
	ErrUnauthorizedNoCredentials      = "Unauthorized. No credentials provided"
)

type Server struct {
	logger *zap.Logger
}

func (s *Server) respondWithError(w http.ResponseWriter, code int, message string) {
	s.respondWithJSON(w, code, map[string]string{"error": message})
}

func (s *Server) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func (s *Server) getUsers(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Inside getUserss")
	name := r.URL.Query().Get("name")
	if name != "" {
		var filteredUsers []User
		for _, user := range users {
			if user.Name == name {
				filteredUsers = append(filteredUsers, user)
			}
		}
		s.respondWithJSON(w, http.StatusOK, filteredUsers)
		return
	}
	s.respondWithJSON(w, http.StatusOK, users)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.respondWithError(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}
	for _, user := range users {
		if user.ID == id {
			s.respondWithJSON(w, http.StatusOK, user)
			return
		}
	}
	s.respondWithError(w, http.StatusNotFound, ErrUserNotFound)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		s.respondWithError(w, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	s.respondWithJSON(w, http.StatusCreated, newUser)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.respondWithError(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		s.respondWithError(w, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			s.respondWithJSON(w, http.StatusOK, users[i])
			return
		}
	}
	s.respondWithError(w, http.StatusNotFound, ErrUserNotFound)
}

func (s *Server) updateOrCreateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.respondWithError(w, http.StatusBadRequest, ErrInvalidUserID)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		s.respondWithError(w, http.StatusBadRequest, ErrInvalidRequestBody)
		return
	}

	for i, existingUser := range users {
		if existingUser.ID == id {
			users[i].Name = user.Name
			s.respondWithJSON(w, http.StatusOK, users[i])
			return
		}
	}

	user.ID = id
	users = append(users, user)
	s.respondWithJSON(w, http.StatusCreated, user)
}

func configureLogger() *zap.Logger {
	// Read the zap configuration from the external JSON file
	configFile, err := os.ReadFile("zap_config.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to read zap config file: %v", err))
	}

	var cfg zap.Config
	if err := json.Unmarshal(configFile, &cfg); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal zap config: %v", err))
	}

	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	logger.Info("logger construction succeeded")

	return logger
}

func authMiddleware(server *Server) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()
			if !ok || username != "admin" || password != "admin" {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)

				errMsg := ErrUnauthorizedInvalidCredentials
				if !ok {
					errMsg = ErrUnauthorizedNoCredentials
				}

				server.respondWithError(w, http.StatusUnauthorized, errMsg)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	logger := configureLogger()

	server := &Server{logger: logger}

	r := chi.NewRouter()

	//middlewares
	r.Use(authMiddleware(server))

	// Routes
	r.Get("/users", server.getUsers)
	r.Get("/users/{id}", server.getUser)
	r.Post("/users", server.createUser)
	r.Patch("/users/{id}", server.updateUser)
	r.Put("/users/{id}", server.updateOrCreateUser)

	// Get the port from the environment variable, default to 3005 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		return
	}
	fmt.Println("Server started on port " + port)
}
