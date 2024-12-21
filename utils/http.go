package utils

import (
	"encoding/json"
	"net/http"

	"github.com/alae-touba/playing-with-go-chi/constants"
)

// RespondWithJSON writes JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set(constants.HeaderContentType, constants.HeaderApplicationJSON)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// RespondWithError writes JSON error response
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
