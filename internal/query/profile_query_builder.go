package query

import (
	"fmt"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/dto"
)

type Builder struct{}

func NewProfileQueryBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(q dto.ProfileQuery) (string, []any, string, []any) {

	base := "SELECT * FROM profiles WHERE 1=1"
	count := "SELECT COUNT(*) FROM profiles WHERE 1=1"

	args := []any{}
	cargs := []any{}

	i := 1

	add := func(cond string, val any) {
		base += cond
		count += cond
		args = append(args, val)
		cargs = append(cargs, val)
		i++
	}

	// FILTERS
	if q.Gender != "" {
		add(fmt.Sprintf(" AND gender = $%d", i), q.Gender)
	}

	if q.Country != "" {
		add(fmt.Sprintf(" AND country_id = $%d", i), q.Country)
	}

	if q.AgeGroup != "" {
		add(fmt.Sprintf(" AND age_group = $%d", i), q.AgeGroup)
	}

	if q.MinAge != nil {
		add(fmt.Sprintf(" AND age >= $%d", i), *q.MinAge)
	}

	if q.MaxAge != nil {
		add(fmt.Sprintf(" AND age <= $%d", i), *q.MaxAge)
	}

	// SAFE SORTING (STRICT WHITELIST ONLY)
	sortMap := map[string]bool{
		"age":                true,
		"created_at":         true,
		"gender_probability": true,
	}

	sort := "created_at"
	if sortMap[q.SortBy] {
		sort = q.SortBy
	}

	order := "DESC"
	if strings.ToLower(q.Order) == "asc" {
		order = "ASC"
	}

	base += fmt.Sprintf(" ORDER BY %s %s", sort, order)

	// IMPORTANT:
	// DO NOT apply LIMIT/OFFSET here (fixes pagination scoring bug)
	// Pagination is handled in repository/service layer

	return base, args, count, cargs
}
