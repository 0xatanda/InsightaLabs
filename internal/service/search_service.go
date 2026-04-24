package service

import (
	"github.com/0xatanda/InsightaLabs/internal/domain"
	"github.com/0xatanda/InsightaLabs/internal/parser"
	"github.com/0xatanda/InsightaLabs/internal/query"
	"github.com/0xatanda/InsightaLabs/internal/repository"
)

type SearchService struct {
	repo    *repository.ProfileRepository
	builder *query.Builder
	parser  *parser.Parser
}

func NewSearchService(
	r *repository.ProfileRepository,
	b *query.Builder,
	p *parser.Parser,
) *SearchService {
	return &SearchService{
		repo:    r,
		builder: b,
		parser:  p,
	}
}

func (s *SearchService) Search(q string, page, limit int) ([]domain.Profile, int, bool, error) {

	filters, ok := s.parser.Parse(q)
	if !ok {
		return nil, 0, false, nil
	}

	filters.Page = page
	filters.Limit = limit

	sql, args, countSQL, cargs := s.builder.Build(filters)

	data, err := s.repo.Fetch(sql, args)
	if err != nil {
		return nil, 0, false, err
	}

	total, err := s.repo.Count(countSQL, cargs)
	if err != nil {
		return nil, 0, false, err
	}

	return data, total, true, nil
}
