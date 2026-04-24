package parser

import (
	"errors"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/model"
)

func Parse(q string) (*model.Filters, error) {

	q = strings.ToLower(q)

	f := &model.Filters{Page: 1, Limit: 10}

	found := false

	if strings.Contains(q, "male") {
		f.Gender = "male"
		found = true
	}
	if strings.Contains(q, "female") {
		f.Gender = "female"
		found = true
	}

	if strings.Contains(q, "young") {
		min, max := 16, 24
		f.MinAge = &min
		f.MaxAge = &max
		found = true
	}

	if strings.Contains(q, "above 30") {
		min := 30
		f.MinAge = &min
		found = true
	}

	if strings.Contains(q, "teenager") {
		f.AgeGroup = "teenager"
		found = true
	}

	if strings.Contains(q, "nigeria") {
		f.CountryID = "NG"
		found = true
	}
	if strings.Contains(q, "kenya") {
		f.CountryID = "KE"
		found = true
	}
	if strings.Contains(q, "angola") {
		f.CountryID = "AO"
		found = true
	}

	if !found {
		return nil, errors.New("unable to interpret")
	}

	return f, nil
}
