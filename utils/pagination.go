package utils

import (
	"net/http"
	"strconv"
)

// pagination
type PaginationParams struct {
	Page    int
	PerPage int
	Offset  int
	Limit   int
}

func GetPaginationParams(r *http.Request) PaginationParams {
	perPage := 10 // default limit
	page := 1     // default page

	if perPageStr := r.URL.Query().Get("per_page"); perPageStr != "" {
		if parsed, err := strconv.Atoi(perPageStr); err == nil && parsed >= 0 {
			perPage = parsed
		}
	}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if parsed, err := strconv.Atoi(pageStr); err == nil && parsed > 0 {
			page = parsed
		}
	}

	offset := (page - 1) * perPage
	limit := perPage

	return PaginationParams{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
		Limit:   limit,
	}
}
