package main

import (
	"database/sql"
	"encoding/json"
	"os"

	"github.com/0xatanda/InsightaLabs/internal/utils"
)

type SeedProfile struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	CountryID string `json:"country_id"`
	Country   string `json:"country_name"`
}

func Seed(db *sql.DB) error {
	file, err := os.ReadFile("scripts/seed.json")
	if err != nil {
		return err
	}

	var data []SeedProfile
	if err := json.Unmarshal(file, &data); err != nil {
		return err
	}

	for _, p := range data {

		_, err := db.Exec(`
			INSERT INTO profiles (
				id, name, gender, age, age_group,
				country_id, country_name, created_at
			)
			VALUES ($1,$2,$3,$4,$5,$6,$7,NOW())
			ON CONFLICT (name) DO NOTHING
		`,
			utils.UUIDv7(),
			p.Name,
			p.Gender,
			p.Age,
			ageGroup(p.Age),
			p.CountryID,
			p.Country,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func ageGroup(age int) string {
	switch {
	case age < 13:
		return "child"
	case age <= 19:
		return "teenager"
	case age <= 59:
		return "adult"
	default:
		return "senior"
	}
}
