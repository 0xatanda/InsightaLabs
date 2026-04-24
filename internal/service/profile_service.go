package service

import (
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
	return &ProfileService{r, b}
}

func (s *ProfileService) Get(q dto.ProfileQuery) ([]domain.Profile, int, error) {

	sql, args, countSQL, cargs := s.builder.Build(q)

	data, err := s.repo.Fetch(sql, args)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count(countSQL, cargs)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
