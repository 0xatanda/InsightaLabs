package parser

import (
	"errors"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/model"
)

func Parse(q string) (*model.Filters, error) {
	q = strings.ToLower(q)

	f := &model.Filters{
		Page:  1,
		Limit: 10,
	}

	// gender
	if strings.Contains(q, "male") {
		f.Gender = "male"
	}
	if strings.Contains(q, "female") {
		f.Gender = "female"
	}

	// age groups
	if strings.Contains(q, "child") {
		f.AgeGroup = "child"
	}
	if strings.Contains(q, "teen") {
		f.AgeGroup = "teenager"
	}
	if strings.Contains(q, "adult") {
		f.AgeGroup = "adult"
	}
	if strings.Contains(q, "senior") {
		f.AgeGroup = "senior"
	}

	// young rule
	if strings.Contains(q, "young") {
		min := 16
		max := 24
		f.MinAge = &min
		f.MaxAge = &max
	}

	// country
	if strings.Contains(q, "nigeria") {
		f.CountryID = "NG"
	}
	if strings.Contains(q, "kenya") {
		f.CountryID = "KE"
	}
	if strings.Contains(q, "angola") {
		f.CountryID = "AO"
	}

	if f.Gender == "" &&
		f.AgeGroup == "" &&
		f.CountryID == "" &&
		f.MinAge == nil &&
		f.MaxAge == nil {
		return nil, errors.New("unable to interpret query")
	}

	return f, nil
}
