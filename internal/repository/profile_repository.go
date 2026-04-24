package repository

import (
	"database/sql"

	"github.com/0xatanda/InsightaLabs/internal/domain"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (r *ProfileRepository) Fetch(query string, args []any) ([]domain.Profile, error) {
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []domain.Profile

	for rows.Next() {
		var p domain.Profile
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}

	return res, nil
}

func (r *ProfileRepository) Count(query string, args []any) (int, error) {
	var c int
	err := r.DB.QueryRow(query, args...).Scan(&c)
	return c, err
}
