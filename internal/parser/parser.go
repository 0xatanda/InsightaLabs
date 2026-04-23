package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/model"
)

var countryMap = map[string]string{
	"nigeria": "NG",
	"kenya":   "KE",
	"angola":  "AO",
	"benin":   "BJ",
	"ghana":   "GH",
}

func Parse(q string) (*model.Filters, error) {
	q = strings.ToLower(q)

	f := &model.Filters{}

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

	// "young" rule (special)
	if strings.Contains(q, "young") {
		min := 16
		max := 24
		f.MinAge = &min
		f.MaxAge = &max
	}

	// "above X"
	if strings.Contains(q, "above") {
		var val int
		_, err := fmt.Sscanf(q, "%*s above %d", &val)
		if err == nil {
			f.MinAge = &val
		}
	}

	// country detection
	for k, v := range countryMap {
		if strings.Contains(q, k) {
			f.CountryID = v
			break
		}
	}

	// validation
	if f.Gender == "" &&
		f.AgeGroup == "" &&
		f.CountryID == "" &&
		f.MinAge == nil &&
		f.MaxAge == nil {
		return nil, errors.New("unable to interpret")
	}

	return f, nil
}
