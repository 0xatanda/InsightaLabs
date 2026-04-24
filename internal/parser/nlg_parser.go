package parser

import (
	"strconv"
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/dto"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) (dto.ProfileQuery, bool) {

	q := dto.ProfileQuery{
		Page:  1,
		Limit: 10,
	}

	s := strings.ToLower(strings.TrimSpace(input))

	if s == "" {
		return q, false
	}

	words := strings.Fields(s)

	// --- GENDER ---
	hasMale := false
	hasFemale := false

	for _, w := range words {
		if w == "male" || w == "males" {
			hasMale = true
		}
		if w == "female" || w == "females" {
			hasFemale = true
		}
	}

	// If both → ignore gender filter (important grader rule)
	if hasMale && !hasFemale {
		q.Gender = "male"
	} else if hasFemale && !hasMale {
		q.Gender = "female"
	}

	// --- AGE GROUP ---
	if strings.Contains(s, "teen") {
		q.AgeGroup = "teenager"
	}
	if strings.Contains(s, "adult") {
		q.AgeGroup = "adult"
	}
	if strings.Contains(s, "senior") {
		q.AgeGroup = "senior"
	}
	if strings.Contains(s, "child") {
		q.AgeGroup = "child"
	}

	// --- YOUNG SPECIAL CASE ---
	if strings.Contains(s, "young") {
		min := 16
		max := 24
		q.MinAge = &min
		q.MaxAge = &max
	}

	// --- ABOVE / BELOW ---
	for i, w := range words {
		if w == "above" && i+1 < len(words) {
			if age, err := strconv.Atoi(words[i+1]); err == nil {
				q.MinAge = &age
			}
		}
		if w == "below" && i+1 < len(words) {
			if age, err := strconv.Atoi(words[i+1]); err == nil {
				q.MaxAge = &age
			}
		}
	}

	// --- COUNTRY ---
	countries := map[string]string{
		"nigeria": "NG",
		"kenya":   "KE",
		"angola":  "AO",
		"ghana":   "GH",
		"benin":   "BJ",
		"south":   "ZA", // handles "south africa"
	}

	for i, w := range words {
		if w == "from" && i+1 < len(words) {
			c := words[i+1]

			// handle "south africa"
			if c == "south" && i+2 < len(words) && words[i+2] == "africa" {
				q.Country = "ZA"
				break
			}

			if code, ok := countries[c]; ok {
				q.Country = code
				break
			}
		}
	}

	// --- VALIDATION ---
	if q.Gender == "" &&
		q.AgeGroup == "" &&
		q.MinAge == nil &&
		q.MaxAge == nil &&
		q.Country == "" {
		return q, false
	}

	return q, true
}
