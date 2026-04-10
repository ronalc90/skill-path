package service

import (
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/repository"
)

// DashboardService handles dashboard aggregation business logic.
type DashboardService struct {
	progressRepo   *repository.ProgressRepository
	assessmentRepo *repository.AssessmentRepository
}

func NewDashboardService(progressRepo *repository.ProgressRepository, assessmentRepo *repository.AssessmentRepository) *DashboardService {
	return &DashboardService{
		progressRepo:   progressRepo,
		assessmentRepo: assessmentRepo,
	}
}

// GetDashboard aggregates all dashboard data for a user.
func (s *DashboardService) GetDashboard(userID uint) (*dto.DashboardResponse, error) {
	progresses, err := s.progressRepo.FindAllUserPaths(userID)
	if err != nil {
		return nil, err
	}

	var pathsInProgress, pathsCompleted int
	var totalHours float64
	enrolledPaths := make([]dto.UserPathProgress, 0)
	recentActivity := make([]dto.RecentActivity, 0)
	categoryScores := make(map[string]float64)
	categoryMax := make(map[string]float64)

	for _, p := range progresses {
		totalMilestones := len(p.Path.Milestones)
		completedCount := 0
		for _, mp := range p.MilestoneProgresses {
			if mp.Status == "completed" {
				completedCount++
			}
		}

		var pct float64
		if totalMilestones > 0 {
			pct = float64(completedCount) / float64(totalMilestones) * 100
		}

		if p.CompletedAt != nil {
			pathsCompleted++
		} else {
			pathsInProgress++
		}

		totalHours += float64(p.Path.EstimatedHours) * (pct / 100)

		milestones := make([]dto.MilestoneProgressDetail, 0)
		for _, mp := range p.MilestoneProgresses {
			milestones = append(milestones, dto.MilestoneProgressDetail{
				MilestoneID: mp.MilestoneID,
				Title:       mp.Milestone.Title,
				Status:      string(mp.Status),
				OrderIndex:  mp.Milestone.OrderIndex,
			})
		}

		enrolledPaths = append(enrolledPaths, dto.UserPathProgress{
			PathID:         p.PathID,
			Path:           toPathSummary(p.Path),
			StartedAt:      p.StartedAt,
			CompletedAt:    p.CompletedAt,
			LastActivityAt: p.LastActivityAt,
			ProgressPct:    pct,
			Milestones:     milestones,
		})

		recentActivity = append(recentActivity, dto.RecentActivity{
			Type:      "progress",
			Title:     p.Path.Title,
			PathTitle: p.Path.Title,
			Timestamp: p.LastActivityAt,
		})

		cat := p.Path.Category
		categoryMax[cat] += 100
		categoryScores[cat] += pct
	}

	assessmentResults, err := s.assessmentRepo.FindUserResults(userID)
	if err == nil {
		for _, r := range assessmentResults {
			milestone, err := s.assessmentRepo.FindMilestoneByID(r.MilestoneID)
			if err != nil {
				continue
			}
			recentActivity = append(recentActivity, dto.RecentActivity{
				Type:      "assessment",
				Title:     milestone.Title,
				PathTitle: "",
				Timestamp: r.CreatedAt,
			})
		}
	}

	skillRadar := make([]dto.SkillRadarPoint, 0)
	for cat, maxScore := range categoryMax {
		skillRadar = append(skillRadar, dto.SkillRadarPoint{
			Category: cat,
			Score:    categoryScores[cat],
			MaxScore: maxScore,
		})
	}

	var completionRate float64
	totalPaths := pathsInProgress + pathsCompleted
	if totalPaths > 0 {
		completionRate = float64(pathsCompleted) / float64(totalPaths) * 100
	}

	if len(recentActivity) > 10 {
		recentActivity = recentActivity[:10]
	}

	return &dto.DashboardResponse{
		PathsInProgress:   pathsInProgress,
		PathsCompleted:    pathsCompleted,
		TotalHoursLearned: totalHours,
		CompletionRate:    completionRate,
		RecentActivity:    recentActivity,
		SkillRadar:        skillRadar,
		EnrolledPaths:     enrolledPaths,
	}, nil
}
