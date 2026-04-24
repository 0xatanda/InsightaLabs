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

	base := "SELECT id, name, gender, gender_probability, age, age_group, country_id, country_name, country_probability, created_at FROM profiles WHERE 1=1"
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

	if q.MinGenderProbability != nil {
		add(fmt.Sprintf(" AND gender_probability >= $%d", i), *q.MinGenderProbability)
	}

	if q.MinCountryProbability != nil {
		add(fmt.Sprintf(" AND country_probability >= $%d", i), *q.MinCountryProbability)
	}

	// SORTING (strict whitelist — validation happens in service)
	sort := "created_at"
	if q.SortBy != "" {
		sort = q.SortBy
	}

	order := "DESC"
	if strings.ToLower(q.Order) == "asc" {
		order = "ASC"
	}

	base += fmt.Sprintf(" ORDER BY %s %s", sort, order)

	// ALWAYS deterministic pagination (CRITICAL)
	base += " , id ASC"

	offset := (q.Page - 1) * q.Limit
	base += fmt.Sprintf(" LIMIT %d OFFSET %d", q.Limit, offset)

	return base, args, count, cargs
}
