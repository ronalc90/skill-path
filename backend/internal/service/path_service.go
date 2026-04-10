package service

import (
	"math"

	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/model"
	"github.com/ronalc90/skillpath/internal/repository"
)

// PathService handles business logic for learning paths.
type PathService struct {
	pathRepo *repository.PathRepository
}

func NewPathService(pathRepo *repository.PathRepository) *PathService {
	return &PathService{pathRepo: pathRepo}
}

// ListPaths returns a paginated list of learning paths with optional filters.
func (s *PathService) ListPaths(category, difficulty, search string, page, pageSize int) (*dto.PathListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 12
	}

	paths, total, err := s.pathRepo.FindAll(category, difficulty, search, page, pageSize)
	if err != nil {
		return nil, err
	}

	summaries := make([]dto.PathSummary, 0, len(paths))
	for _, p := range paths {
		summaries = append(summaries, toPathSummary(p))
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &dto.PathListResponse{
		Paths:      summaries,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetPathBySlug returns a full path detail including milestones and resources.
func (s *PathService) GetPathBySlug(slug string) (*dto.PathDetailResponse, error) {
	path, err := s.pathRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	milestones := make([]dto.MilestoneResponse, 0, len(path.Milestones))
	for _, m := range path.Milestones {
		resources := make([]dto.ResourceResponse, 0, len(m.Resources))
		for _, r := range m.Resources {
			resources = append(resources, dto.ResourceResponse{
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
		milestones = append(milestones, dto.MilestoneResponse{
			ID:             m.ID,
			Title:          m.Title,
			Description:    m.Description,
			OrderIndex:     m.OrderIndex,
			EstimatedHours: m.EstimatedHours,
			Resources:      resources,
		})
	}

	return &dto.PathDetailResponse{
		ID:             path.ID,
		Title:          path.Title,
		Slug:           path.Slug,
		Description:    path.Description,
		Difficulty:     string(path.Difficulty),
		EstimatedHours: path.EstimatedHours,
		Category:       path.Category,
		IconURL:        path.IconURL,
		IsPublished:    path.IsPublished,
		Milestones:     milestones,
	}, nil
}

// GetCategories returns all categories with their path counts.
func (s *PathService) GetCategories() ([]dto.CategoryResponse, error) {
	results, err := s.pathRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	categories := make([]dto.CategoryResponse, 0, len(results))
	for _, r := range results {
		categories = append(categories, dto.CategoryResponse{
			Name:  r.Category,
			Count: r.Count,
		})
	}

	return categories, nil
}

func toPathSummary(p model.LearningPath) dto.PathSummary {
	return dto.PathSummary{
		ID:             p.ID,
		Title:          p.Title,
		Slug:           p.Slug,
		Description:    p.Description,
		Difficulty:     string(p.Difficulty),
		EstimatedHours: p.EstimatedHours,
		Category:       p.Category,
		IconURL:        p.IconURL,
		MilestoneCount: len(p.Milestones),
	}
}
