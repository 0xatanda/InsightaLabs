package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/external"
	"github.com/0xatanda/InsightaLabs/internal/model"
	"github.com/0xatanda/InsightaLabs/internal/utils"
)

type Repo struct {
	DB *sql.DB
}

func (r *Repo) CreateOrGetProfile(name string) (*model.Profile, bool, error) {

	var p model.Profile

	err := r.DB.QueryRow(`
	SELECT id,name,gender,gender_probability,age,age_group,
	country_id,country_name,country_probability,created_at
	FROM profiles WHERE LOWER(name)=LOWER($1)
	`, name).Scan(
		&p.ID, &p.Name, &p.Gender, &p.GenderProbability,
		&p.Age, &p.AgeGroup, &p.CountryID, &p.CountryName,
		&p.CountryProbability, &p.CreatedAt,
	)

	if err == nil {
		return &p, true, nil
	}

	g, a, c, err := external.FetchAll(name)
	if err != nil {
		return nil, false, fmt.Errorf("%s returned invalid response", err)
	}

	var ageGroup string
	switch {
	case a.Age <= 12:
		ageGroup = "child"
	case a.Age <= 19:
		ageGroup = "teenager"
	case a.Age <= 59:
		ageGroup = "adult"
	default:
		ageGroup = "senior"
	}

	countryName := utils.CountryMap[c.CountryID]

	err = r.DB.QueryRow(`
	INSERT INTO profiles (
	name,gender,gender_probability,age,age_group,
	country_id,country_name,country_probability
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	RETURNING id,created_at
	`,
		name, g.Gender, g.Prob,
		a.Age, ageGroup,
		c.CountryID, countryName, c.Prob,
	).Scan(&p.ID, &p.CreatedAt)

	if err != nil {
		return nil, false, err
	}

	p = model.Profile{
		ID:                 p.ID,
		Name:               name,
		Gender:             g.Gender,
		GenderProbability:  g.Prob,
		Age:                a.Age,
		AgeGroup:           ageGroup,
		CountryID:          c.CountryID,
		CountryName:        countryName,
		CountryProbability: c.Prob,
		CreatedAt:          p.CreatedAt,
	}

	return &p, false, nil
}

func (r *Repo) FindAll(f *model.Filters) ([]model.Profile, int, error) {

	query := `SELECT id,name,gender,gender_probability,age,age_group,
	country_id,country_name,country_probability,created_at FROM profiles WHERE 1=1`

	args := []interface{}{}
	i := 1

	if f.Gender != "" {
		query += fmt.Sprintf(" AND LOWER(gender)=LOWER($%d)", i)
		args = append(args, f.Gender)
		i++
	}

	if f.CountryID != "" {
		query += fmt.Sprintf(" AND country_id=$%d", i)
		args = append(args, f.CountryID)
		i++
	}

	if f.MinAge != nil {
		query += fmt.Sprintf(" AND age >= $%d", i)
		args = append(args, *f.MinAge)
		i++
	}

	if f.MaxAge != nil {
		query += fmt.Sprintf(" AND age <= $%d", i)
		args = append(args, *f.MaxAge)
		i++
	}

	if f.SortBy != "" {
		allowed := map[string]bool{"age": true, "created_at": true, "gender_probability": true}
		if allowed[f.SortBy] {
			order := "ASC"
			if strings.ToLower(f.Order) == "desc" {
				order = "DESC"
			}
			query += fmt.Sprintf(" ORDER BY %s %s", f.SortBy, order)
		}
	}

	offset := (f.Page - 1) * f.Limit
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", f.Limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []model.Profile

	for rows.Next() {
		var p model.Profile
		_ = rows.Scan(
			&p.ID, &p.Name, &p.Gender, &p.GenderProbability,
			&p.Age, &p.AgeGroup, &p.CountryID, &p.CountryName,
			&p.CountryProbability, &p.CreatedAt,
		)
		list = append(list, p)
	}

	var total int
	_ = r.DB.QueryRow("SELECT COUNT(*) FROM profiles").Scan(&total)

	return list, total, nil
}
