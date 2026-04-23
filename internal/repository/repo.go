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

	base := "FROM profiles WHERE 1=1"
	args := []any{}
	i := 1

	if f.Gender != "" {
		base += fmt.Sprintf(" AND gender=$%d", i)
		args = append(args, f.Gender)
		i++
	}

	if f.AgeGroup != "" {
		base += fmt.Sprintf(" AND age_group=$%d", i)
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

	// COUNT
	var total int
	countQ := "SELECT COUNT(*) " + base
	if err := r.DB.QueryRow(countQ, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// SORT
	sortBy := "created_at"
	order := "DESC"

	if f.SortBy != "" {
		allowed := map[string]bool{
			"age":                true,
			"created_at":         true,
			"gender_probability": true,
		}
		if allowed[f.SortBy] {
			sortBy = f.SortBy
		}
	}

	if strings.ToLower(f.Order) == "asc" {
		order = "ASC"
	}

	// QUERY
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

	var res []model.Profile

	for rows.Next() {
		var p model.Profile
		rows.Scan(
			&p.ID, &p.Name, &p.Gender,
			&p.GenderProbability, &p.Age,
			&p.AgeGroup, &p.CountryID,
			&p.CountryName, &p.CountryProbability,
			&p.CreatedAt,
		)
		res = append(res, p)
	}

	return res, total, nil
}

func (r *Repo) Create(p *model.Profile) error {
	query := `
	INSERT INTO profiles (
		id, name, gender, gender_probability,
		age, age_group, country_id, country_name,
		country_probability, created_at
	) VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10
	)
	ON CONFLICT (name) DO NOTHING
	`

	_, err := r.DB.Exec(
		query,
		p.ID,
		p.Name,
		p.Gender,
		p.GenderProbability,
		p.Age,
		p.AgeGroup,
		p.CountryID,
		p.CountryName,
		p.CountryProbability,
		p.CreatedAt,
	)

	return err
}
