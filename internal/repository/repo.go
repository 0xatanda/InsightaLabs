package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/model"
)

type Repo struct {
	DB *sql.DB
}

func (r *Repo) FindAll(f *model.Filters) ([]model.Profile, int, error) {

	base := `FROM profiles WHERE 1=1`
	args := []interface{}{}
	i := 1

	if f.Gender != "" {
		base += fmt.Sprintf(" AND LOWER(gender)=LOWER($%d)", i)
		args = append(args, f.Gender)
		i++
	}

	if f.AgeGroup != "" {
		base += fmt.Sprintf(" AND LOWER(age_group)=LOWER($%d)", i)
		args = append(args, f.AgeGroup)
		i++
	}

	if f.CountryID != "" {
		base += fmt.Sprintf(" AND country_id=$%d", i)
		args = append(args, f.CountryID)
		i++
	}

	if f.MinAge != nil {
		base += fmt.Sprintf(" AND age >= $%d", i)
		args = append(args, *f.MinAge)
		i++
	}

	if f.MaxAge != nil {
		base += fmt.Sprintf(" AND age <= $%d", i)
		args = append(args, *f.MaxAge)
		i++
	}

	if f.MinGenderProbability != nil {
		base += fmt.Sprintf(" AND gender_probability >= $%d", i)
		args = append(args, *f.MinGenderProbability)
		i++
	}

	if f.MinCountryProbability != nil {
		base += fmt.Sprintf(" AND country_probability >= $%d", i)
		args = append(args, *f.MinCountryProbability)
		i++
	}

	// total count (FILTERED)
	var total int
	countQuery := "SELECT COUNT(*) " + base
	if err := r.DB.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// sorting
	sortBy := "created_at"
	order := "DESC"

	switch f.SortBy {
	case "age", "created_at", "gender_probability":
		sortBy = f.SortBy
	}

	if strings.ToLower(f.Order) == "asc" {
		order = "ASC"
	}

	query := `
	SELECT id, name, gender, gender_probability,
	       age, age_group, country_id, country_name,
	       country_probability, created_at
	` + base + fmt.Sprintf(" ORDER BY %s %s", sortBy, order)

	offset := (f.Page - 1) * f.Limit
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, f.Limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var profiles []model.Profile

	for rows.Next() {
		var p model.Profile
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Gender,
			&p.GenderProbability,
			&p.Age,
			&p.AgeGroup,
			&p.CountryID,
			&p.CountryName,
			&p.CountryProbability,
			&p.CreatedAt,
		); err != nil {
			return nil, 0, err
		}
		profiles = append(profiles, p)
	}

	return profiles, total, nil
}
