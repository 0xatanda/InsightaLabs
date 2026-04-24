package parser

import (
	"strings"

	"github.com/0xatanda/InsightaLabs/internal/dto"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(q string) (dto.ProfileQuery, bool) {

	q = strings.ToLower(q)

	var query dto.ProfileQuery
	found := false

	if strings.Contains(q, "male") {
		query.Gender = "male"
		found = true
	}

	if strings.Contains(q, "female") {
		query.Gender = "female"
		found = true
	}

	if strings.Contains(q, "young") {
		min, max := 16, 24
		query.MinAge = &min
		query.MaxAge = &max
		found = true
	}

	if strings.Contains(q, "adult") {
		query.AgeGroup = "adult"
		found = true
	}

	if strings.Contains(q, "nigeria") {
		query.Country = "NG"
		found = true
	}

	if strings.Contains(q, "kenya") {
		query.Country = "KE"
		found = true
	}

	return query, found
}
