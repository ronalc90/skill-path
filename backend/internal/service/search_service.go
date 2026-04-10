package service

import (
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/repository"
)

// SearchService handles search business logic.
type SearchService struct {
	pathRepo *repository.PathRepository
}

func NewSearchService(pathRepo *repository.PathRepository) *SearchService {
	return &SearchService{pathRepo: pathRepo}
}

// Search searches paths and resources by keyword.
func (s *SearchService) Search(query string) (*dto.SearchResponse, error) {
	if query == "" {
		return &dto.SearchResponse{
			Paths:     []dto.PathSummary{},
			Resources: []dto.ResourceResponse{},
		}, nil
	}

	paths, err := s.pathRepo.SearchPaths(query)
	if err != nil {
		return nil, err
	}

	resources, err := s.pathRepo.SearchResources(query)
	if err != nil {
		return nil, err
	}

	pathSummaries := make([]dto.PathSummary, 0, len(paths))
	for _, p := range paths {
		pathSummaries = append(pathSummaries, toPathSummary(p))
	}

	resourceResponses := make([]dto.ResourceResponse, 0, len(resources))
	for _, r := range resources {
		resourceResponses = append(resourceResponses, dto.ResourceResponse{
			ID:               r.ID,
			Title:            r.Title,
			URL:              r.URL,
			Type:             string(r.Type),
			Provider:         r.Provider,
			IsFree:           r.IsFree,
			EstimatedMinutes: r.EstimatedMinutes,
			OrderIndex:       r.OrderIndex,
		})
	}

	return &dto.SearchResponse{
		Paths:     pathSummaries,
		Resources: resourceResponses,
	}, nil
}
