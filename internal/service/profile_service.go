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

func (s *ProfileService) Get(q dto.ProfileQuery) ([]domain.Profile, int, error) {

	// ✅ VALIDATE sort_by
	validSort := map[string]bool{
		"age":                true,
		"created_at":         true,
		"gender_probability": true,
		"":                   true, // allow empty → default
	}

	if !validSort[q.SortBy] {
		return nil, 0, fmt.Errorf("Invalid query parameters")
	}

	// ✅ VALIDATE order
	if q.Order != "" && q.Order != "asc" && q.Order != "desc" {
		return nil, 0, fmt.Errorf("Invalid query parameters")
	}

	query, args, countQuery, countArgs := s.builder.Build(q)

	data, err := s.repo.Fetch(query, args)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count(countQuery, countArgs)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
