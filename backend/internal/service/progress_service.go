package service

import (
	"errors"
	"time"

	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/model"
	"github.com/ronalc90/skillpath/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrAlreadyEnrolled         = errors.New("already enrolled in this path")
	ErrNotEnrolled             = errors.New("not enrolled in this path")
	ErrPathNotFound            = errors.New("learning path not found")
	ErrResourcesNotCompleted   = errors.New("all resources must be completed before completing the milestone")
)

// ProgressService handles progress tracking business logic.
type ProgressService struct {
	progressRepo *repository.ProgressRepository
	pathRepo     *repository.PathRepository
}

func NewProgressService(progressRepo *repository.ProgressRepository, pathRepo *repository.PathRepository) *ProgressService {
	return &ProgressService{
		progressRepo: progressRepo,
		pathRepo:     pathRepo,
	}
}

// FindPathIDBySlug resolves a slug to a path ID.
func (s *ProgressService) FindPathIDBySlug(slug string) (uint, error) {
	path, err := s.pathRepo.FindBySlug(slug)
	if err != nil {
		return 0, ErrPathNotFound
	}
	return path.ID, nil
}

// StartPath enrolls a user in a learning path.
func (s *ProgressService) StartPath(userID, pathID uint) (*dto.UserPathProgress, error) {
	path, err := s.pathRepo.FindByID(pathID)
	if err != nil {
		return nil, ErrPathNotFound
	}

	existing, _ := s.progressRepo.FindUserProgress(userID, pathID)
	if existing != nil {
		return nil, ErrAlreadyEnrolled
	}

	now := time.Now()
	progress := &model.UserProgress{
		UserID:         userID,
		PathID:         pathID,
		StartedAt:      now,
		LastActivityAt: now,
	}

	if err := s.progressRepo.CreateUserProgress(progress); err != nil {
		return nil, err
	}

	return &dto.UserPathProgress{
		PathID: pathID,
		Path: dto.PathSummary{
			ID:             path.ID,
			Title:          path.Title,
			Slug:           path.Slug,
			Description:    path.Description,
			Difficulty:     string(path.Difficulty),
			EstimatedHours: path.EstimatedHours,
			Category:       path.Category,
			MilestoneCount: len(path.Milestones),
		},
		StartedAt:      now,
		LastActivityAt: now,
		ProgressPct:    0,
	}, nil
}

// GetPathProgress returns the user's progress on a specific path.
func (s *ProgressService) GetPathProgress(userID, pathID uint) (*dto.UserPathProgress, error) {
	progress, err := s.progressRepo.FindUserProgress(userID, pathID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotEnrolled
		}
		return nil, err
	}

	path, err := s.pathRepo.FindByID(pathID)
	if err != nil {
		return nil, ErrPathNotFound
	}

	pct := s.calculateProgress(progress, path)

	milestones := make([]dto.MilestoneProgressDetail, 0)
	for _, mp := range progress.MilestoneProgresses {
		milestones = append(milestones, dto.MilestoneProgressDetail{
			MilestoneID: mp.MilestoneID,
			Title:       mp.Milestone.Title,
			Status:      string(mp.Status),
			OrderIndex:  mp.Milestone.OrderIndex,
		})
	}

	return &dto.UserPathProgress{
		PathID:         pathID,
		Path:           toPathSummary(*path),
		StartedAt:      progress.StartedAt,
		CompletedAt:    progress.CompletedAt,
		LastActivityAt: progress.LastActivityAt,
		ProgressPct:    pct,
		Milestones:     milestones,
	}, nil
}

// CompleteMilestone marks a milestone as completed for a user.
// All resources in the milestone must be completed first.
func (s *ProgressService) CompleteMilestone(userID, pathID, milestoneID uint) error {
	progress, err := s.progressRepo.FindUserProgress(userID, pathID)
	if err != nil {
		return ErrNotEnrolled
	}

	// Validate that all resources in the milestone are completed
	totalResources, err := s.progressRepo.CountResourcesForMilestone(milestoneID)
	if err != nil {
		return err
	}
	if totalResources > 0 {
		completedResources, err := s.progressRepo.CountCompletedResourcesForMilestone(progress.ID, milestoneID)
		if err != nil {
			return err
		}
		if completedResources < totalResources {
			return ErrResourcesNotCompleted
		}
	}

	now := time.Now()

	mp, err := s.progressRepo.FindMilestoneProgress(progress.ID, milestoneID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			mp = &model.MilestoneProgress{
				UserProgressID: progress.ID,
				MilestoneID:    milestoneID,
				Status:         model.StatusCompleted,
				CompletedAt:    &now,
			}
			if err := s.progressRepo.CreateMilestoneProgress(mp); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		mp.Status = model.StatusCompleted
		mp.CompletedAt = &now
		if err := s.progressRepo.UpdateMilestoneProgress(mp); err != nil {
			return err
		}
	}

	progress.LastActivityAt = now

	path, err := s.pathRepo.FindByID(pathID)
	if err == nil {
		completedCount, _ := s.progressRepo.CountCompletedMilestones(progress.ID)
		if int(completedCount) >= len(path.Milestones) {
			progress.CompletedAt = &now
		}
	}

	return s.progressRepo.UpdateUserProgress(progress)
}

// CompleteResource marks a resource as completed for a user.
func (s *ProgressService) CompleteResource(userID, pathID, resourceID uint, rating int) error {
	progress, err := s.progressRepo.FindUserProgress(userID, pathID)
	if err != nil {
		return ErrNotEnrolled
	}

	rp, err := s.progressRepo.FindResourceProgress(progress.ID, resourceID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			now := time.Now()
			rp = &model.ResourceProgress{
				UserProgressID: progress.ID,
				ResourceID:     resourceID,
				IsCompleted:    true,
				CompletedAt:    &now,
				Rating:         rating,
			}
			if err := s.progressRepo.CreateResourceProgress(rp); err != nil {
				return err
			}
			progress.LastActivityAt = now
			return s.progressRepo.UpdateUserProgress(progress)
		}
		return err
	}

	now := time.Now()
	rp.IsCompleted = true
	rp.CompletedAt = &now
	if rating > 0 {
		rp.Rating = rating
	}

	if err := s.progressRepo.UpdateResourceProgress(rp); err != nil {
		return err
	}

	progress.LastActivityAt = now
	return s.progressRepo.UpdateUserProgress(progress)
}

// GetMyPaths returns all paths a user is enrolled in with progress.
func (s *ProgressService) GetMyPaths(userID uint) ([]dto.UserPathProgress, error) {
	progresses, err := s.progressRepo.FindAllUserPaths(userID)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UserPathProgress, 0, len(progresses))
	for _, p := range progresses {
		pct := s.calculateProgress(&p, &p.Path)

		milestones := make([]dto.MilestoneProgressDetail, 0)
		for _, mp := range p.MilestoneProgresses {
			milestones = append(milestones, dto.MilestoneProgressDetail{
				MilestoneID: mp.MilestoneID,
				Title:       mp.Milestone.Title,
				Status:      string(mp.Status),
				OrderIndex:  mp.Milestone.OrderIndex,
			})
		}

		result = append(result, dto.UserPathProgress{
			PathID:         p.PathID,
			Path:           toPathSummary(p.Path),
			StartedAt:      p.StartedAt,
			CompletedAt:    p.CompletedAt,
			LastActivityAt: p.LastActivityAt,
			ProgressPct:    pct,
			Milestones:     milestones,
		})
	}

	return result, nil
}

func (s *ProgressService) calculateProgress(progress *model.UserProgress, path *model.LearningPath) float64 {
	totalMilestones := len(path.Milestones)
	if totalMilestones == 0 {
		return 0
	}

	completedCount := 0
	for _, mp := range progress.MilestoneProgresses {
		if mp.Status == model.StatusCompleted {
			completedCount++
		}
	}

	return float64(completedCount) / float64(totalMilestones) * 100
}
