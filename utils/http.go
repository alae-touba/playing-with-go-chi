package utils

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/alae-touba/playing-with-go-chi/constants"
)

// Handle http responses
type ListResponse struct {
	Data    interface{} `json:"data"`
	Count   int         `json:"count"`
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
	Total   int         `json:"total"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set(constants.HeaderContentType, constants.HeaderApplicationJSON)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithList(w http.ResponseWriter, code int, data interface{}, page, perPage, total int) {
	count := 0
	if reflect.TypeOf(data).Kind() == reflect.Slice {
		count = reflect.ValueOf(data).Len()
	}

	listResponse := ListResponse{
		Data:    data,
		Count:   count,
		Page:    page,
		PerPage: perPage,
		Total:   total,
	}
	RespondWithJSON(w, code, listResponse)
}
