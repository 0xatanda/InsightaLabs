package utils

import (
	"net/http"
	"strconv"

	"github.com/0xatanda/InsightaLabs/internal/dto"
)

func ParseRequest(r *http.Request) dto.ProfileQuery {

	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(q.Get("limit"))
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	var minAge, maxAge *int

	if v := q.Get("min_age"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			minAge = &i
		}
	}

	if v := q.Get("max_age"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			maxAge = &i
		}
	}

	return dto.ProfileQuery{
		Gender:   q.Get("gender"),
		Country:  q.Get("country_id"),
		AgeGroup: q.Get("age_group"),
		MinAge:   minAge,
		MaxAge:   maxAge,
		SortBy:   q.Get("sort_by"),
		Order:    q.Get("order"),
		Page:     page,
		Limit:    limit,
	}
}
