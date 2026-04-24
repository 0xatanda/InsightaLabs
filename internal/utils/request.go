package utils

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/dto"
)

func ParseRequest(r *http.Request) dto.ProfileQuery {

	q := r.URL.Query()

	// -----------------------
	// PAGINATION (STRICT FIX)
	// -----------------------
	page, err := strconv.Atoi(q.Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	// -----------------------
	// FILTERS
	// -----------------------
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

	// -----------------------
	// NORMALIZATION (IMPORTANT FOR NLP TESTS)
	// -----------------------
	gender := strings.ToLower(q.Get("gender"))
	if gender == "" {
		gender = q.Get("gender")
	}

	country := strings.ToUpper(q.Get("country_id"))

	sortBy := q.Get("sort_by")
	order := strings.ToUpper(q.Get("order"))
	if order != "ASC" {
		order = "DESC"
	}

	return dto.ProfileQuery{
		Gender:   gender,
		Country:  country,
		AgeGroup: q.Get("age_group"),
		MinAge:   minAge,
		MaxAge:   maxAge,
		SortBy:   sortBy,
		Order:    order,
		Page:     page,
		Limit:    limit,
	}
}
