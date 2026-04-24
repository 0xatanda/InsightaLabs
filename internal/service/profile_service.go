package service

import (
	"fmt"

	"github.com/0xatanda/InsightaLabs/internal/domain"
	"github.com/0xatanda/InsightaLabs/internal/dto"
	"github.com/0xatanda/InsightaLabs/internal/query"
	"github.com/0xatanda/InsightaLabs/internal/repository"
)

type ProfileService struct {
	repo    *repository.ProfileRepository
	builder *query.Builder
}

func NewProfileService(r *repository.ProfileRepository, b *query.Builder) *ProfileService {
	return &ProfileService{
		repo:    r,
		builder: b,
	}
}

var validSort = map[string]bool{
	"":                   true,
	"age":                true,
	"created_at":         true,
	"gender_probability": true,
}

var validOrder = map[string]bool{
	"":     true,
	"asc":  true,
	"desc": true,
}

func (s *ProfileService) Get(q dto.ProfileQuery) ([]domain.Profile, int, error) {

	// -------------------------
	// VALIDATE SORT_BY (PUT HERE)
	// -------------------------
	validSort := map[string]bool{
		"age":                true,
		"created_at":         true,
		"gender_probability": true,
	}

	if q.SortBy != "" && !validSort[q.SortBy] {
		return nil, 0, fmt.Errorf("invalid query parameters")
	}

	// build query
	sqlQuery, args, countQuery, countArgs := s.builder.Build(q)

	data, err := s.repo.Fetch(sqlQuery, args)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count(countQuery, countArgs)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
