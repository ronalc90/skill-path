package service_test

import (
	"testing"

	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/model"
	"github.com/ronalc90/skillpath/internal/repository"
	"github.com/ronalc90/skillpath/internal/service"
	"github.com/ronalc90/skillpath/pkg/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupServiceTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.LearningPath{},
		&model.Milestone{},
		&model.Resource{},
		&model.UserProgress{},
		&model.MilestoneProgress{},
		&model.ResourceProgress{},
		&model.SkillAssessment{},
		&model.AssessmentResult{},
	)
	require.NoError(t, err)

	return db
}

func cleanupServiceTestDB(t *testing.T, db *gorm.DB) {
	t.Helper()
	db.Exec("DELETE FROM assessment_results")
	db.Exec("DELETE FROM skill_assessments")
	db.Exec("DELETE FROM resource_progresses")
	db.Exec("DELETE FROM milestone_progresses")
	db.Exec("DELETE FROM user_progresses")
	db.Exec("DELETE FROM resources")
	db.Exec("DELETE FROM milestones")
	db.Exec("DELETE FROM learning_paths")
	db.Exec("DELETE FROM users")
}

func TestProgressCalculation_ZeroMilestones(t *testing.T) {
	db := setupServiceTestDB(t)
	defer cleanupServiceTestDB(t, db)

	hashed, _ := hash.HashPassword("test123")
	user := &model.User{Email: "calc-test@example.com", PasswordHash: hashed, DisplayName: "Calc User"}
	db.Create(user)

	path := &model.LearningPath{
		Title:       "Empty Path",
		Description: "No milestones",
		Slug:        "empty-path-calc",
		Difficulty:  model.DifficultyBeginner,
		Category:    "Test",
		IsPublished: true,
	}
	db.Create(path)

	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)

	result, err := progressService.StartPath(user.ID, path.ID)
	require.NoError(t, err)
	assert.Equal(t, float64(0), result.ProgressPct)
}

func TestProgressCalculation_HalfCompletion(t *testing.T) {
	db := setupServiceTestDB(t)
	defer cleanupServiceTestDB(t, db)

	hashed, _ := hash.HashPassword("test123")
	user := &model.User{Email: "half-calc@example.com", PasswordHash: hashed, DisplayName: "Half Calc"}
	db.Create(user)

	path := &model.LearningPath{
		Title:       "Half Path",
		Description: "Two milestones",
		Slug:        "half-path-calc",
		Difficulty:  model.DifficultyBeginner,
		Category:    "Test",
		IsPublished: true,
		Milestones: []model.Milestone{
			{Title: "M1", OrderIndex: 1, EstimatedHours: 5},
			{Title: "M2", OrderIndex: 2, EstimatedHours: 5},
		},
	}
	db.Create(path)

	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)

	_, err := progressService.StartPath(user.ID, path.ID)
	require.NoError(t, err)

	err = progressService.CompleteMilestone(user.ID, path.ID, path.Milestones[0].ID)
	require.NoError(t, err)

	progress, err := progressService.GetPathProgress(user.ID, path.ID)
	require.NoError(t, err)
	assert.Equal(t, float64(50), progress.ProgressPct)
}

func TestProgressCalculation_FullCompletion(t *testing.T) {
	db := setupServiceTestDB(t)
	defer cleanupServiceTestDB(t, db)

	hashed, _ := hash.HashPassword("test123")
	user := &model.User{Email: "full-calc@example.com", PasswordHash: hashed, DisplayName: "Full Calc"}
	db.Create(user)

	path := &model.LearningPath{
		Title:       "Full Path",
		Description: "Two milestones",
		Slug:        "full-path-calc",
		Difficulty:  model.DifficultyBeginner,
		Category:    "Test",
		IsPublished: true,
		Milestones: []model.Milestone{
			{Title: "M1", OrderIndex: 1, EstimatedHours: 5},
			{Title: "M2", OrderIndex: 2, EstimatedHours: 5},
		},
	}
	db.Create(path)

	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)

	_, err := progressService.StartPath(user.ID, path.ID)
	require.NoError(t, err)

	err = progressService.CompleteMilestone(user.ID, path.ID, path.Milestones[0].ID)
	require.NoError(t, err)
	err = progressService.CompleteMilestone(user.ID, path.ID, path.Milestones[1].ID)
	require.NoError(t, err)

	progress, err := progressService.GetPathProgress(user.ID, path.ID)
	require.NoError(t, err)
	assert.Equal(t, float64(100), progress.ProgressPct)
	assert.NotNil(t, progress.CompletedAt)
}

func TestAssessmentScoring_PassingScore(t *testing.T) {
	db := setupServiceTestDB(t)
	defer cleanupServiceTestDB(t, db)

	hashed, _ := hash.HashPassword("test123")
	user := &model.User{Email: "score-test@example.com", PasswordHash: hashed, DisplayName: "Score User"}
	db.Create(user)

	path := &model.LearningPath{
		Title:       "Score Path",
		Description: "Score test",
		Slug:        "score-path",
		Difficulty:  model.DifficultyBeginner,
		Category:    "Test",
		IsPublished: true,
		Milestones: []model.Milestone{
			{Title: "Score Milestone", OrderIndex: 1, EstimatedHours: 5},
		},
	}
	db.Create(path)

	assessments := []model.SkillAssessment{
		{MilestoneID: path.Milestones[0].ID, Question: "Q1?", Options: model.JSONSlice{"A", "B", "C"}, CorrectAnswer: 0, Explanation: "A is correct"},
		{MilestoneID: path.Milestones[0].ID, Question: "Q2?", Options: model.JSONSlice{"A", "B", "C"}, CorrectAnswer: 1, Explanation: "B is correct"},
		{MilestoneID: path.Milestones[0].ID, Question: "Q3?", Options: model.JSONSlice{"A", "B", "C"}, CorrectAnswer: 2, Explanation: "C is correct"},
	}
	for i := range assessments {
		db.Create(&assessments[i])
	}

	assessmentRepo := repository.NewAssessmentRepository(db)
	assessmentService := service.NewAssessmentService(assessmentRepo)

	result, err := assessmentService.SubmitAssessment(user.ID, path.Milestones[0].ID, dto.SubmitAssessmentRequest{
		Answers: []dto.AnswerSubmission{
			{QuestionID: assessments[0].ID, SelectedIndex: 0},
			{QuestionID: assessments[1].ID, SelectedIndex: 1},
			{QuestionID: assessments[2].ID, SelectedIndex: 2},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 3, result.Score)
	assert.Equal(t, 3, result.TotalQuestions)
	assert.True(t, result.Passed)
}

func TestAssessmentScoring_FailingScore(t *testing.T) {
	db := setupServiceTestDB(t)
	defer cleanupServiceTestDB(t, db)

	hashed, _ := hash.HashPassword("test123")
	user := &model.User{Email: "fail-test@example.com", PasswordHash: hashed, DisplayName: "Fail User"}
	db.Create(user)

	path := &model.LearningPath{
		Title:       "Fail Path",
		Description: "Fail test",
		Slug:        "fail-path",
		Difficulty:  model.DifficultyBeginner,
		Category:    "Test",
		IsPublished: true,
		Milestones: []model.Milestone{
			{Title: "Fail Milestone", OrderIndex: 1, EstimatedHours: 5},
		},
	}
	db.Create(path)

	assessments := []model.SkillAssessment{
		{MilestoneID: path.Milestones[0].ID, Question: "Q1?", Options: model.JSONSlice{"A", "B", "C"}, CorrectAnswer: 0, Explanation: "A"},
		{MilestoneID: path.Milestones[0].ID, Question: "Q2?", Options: model.JSONSlice{"A", "B", "C"}, CorrectAnswer: 1, Explanation: "B"},
		{MilestoneID: path.Milestones[0].ID, Question: "Q3?", Options: model.JSONSlice{"A", "B", "C"}, CorrectAnswer: 2, Explanation: "C"},
	}
	for i := range assessments {
		db.Create(&assessments[i])
	}

	assessmentRepo := repository.NewAssessmentRepository(db)
	assessmentService := service.NewAssessmentService(assessmentRepo)

	// 1/3 correct = 33% < 70% passing
	result, err := assessmentService.SubmitAssessment(user.ID, path.Milestones[0].ID, dto.SubmitAssessmentRequest{
		Answers: []dto.AnswerSubmission{
			{QuestionID: assessments[0].ID, SelectedIndex: 0}, // correct
			{QuestionID: assessments[1].ID, SelectedIndex: 0}, // wrong
			{QuestionID: assessments[2].ID, SelectedIndex: 0}, // wrong
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 1, result.Score)
	assert.False(t, result.Passed)
}
