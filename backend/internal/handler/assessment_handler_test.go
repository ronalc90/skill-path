package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/handler"
	"github.com/ronalc90/skillpath/internal/model"
	"github.com/ronalc90/skillpath/internal/repository"
	"github.com/ronalc90/skillpath/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupAssessmentHandler(db *gorm.DB) (*handler.AssessmentHandler, *gin.Engine, *model.User) {
	gin.SetMode(gin.TestMode)

	user := &model.User{
		Email:        "assessment-test@example.com",
		PasswordHash: "hashed",
		DisplayName:  "Assessment Test User",
	}
	db.Create(user)

	assessmentRepo := repository.NewAssessmentRepository(db)
	assessmentService := service.NewAssessmentService(assessmentRepo)
	assessmentHandler := handler.NewAssessmentHandler(assessmentService)

	authMiddleware := func(c *gin.Context) {
		c.Set("userID", user.ID)
		c.Next()
	}

	r := gin.New()
	r.GET("/api/v1/milestones/:id/assessment", authMiddleware, assessmentHandler.GetAssessment)
	r.POST("/api/v1/milestones/:id/assessment/submit", authMiddleware, assessmentHandler.SubmitAssessment)

	return assessmentHandler, r, user
}

func seedTestAssessment(t *testing.T, db *gorm.DB, milestoneID uint) []model.SkillAssessment {
	t.Helper()
	assessments := []model.SkillAssessment{
		{
			MilestoneID:   milestoneID,
			Question:      "What is Go?",
			Options:       model.JSONSlice{"A programming language", "A game", "A database", "A framework"},
			CorrectAnswer: 0,
			Explanation:   "Go is a programming language developed by Google.",
		},
		{
			MilestoneID:   milestoneID,
			Question:      "What does := mean in Go?",
			Options:       model.JSONSlice{"Assignment", "Short variable declaration", "Comparison", "Type assertion"},
			CorrectAnswer: 1,
			Explanation:   ":= is used for short variable declaration in Go.",
		},
		{
			MilestoneID:   milestoneID,
			Question:      "What is a goroutine?",
			Options:       model.JSONSlice{"A thread", "A lightweight thread managed by Go runtime", "A function", "A module"},
			CorrectAnswer: 1,
			Explanation:   "A goroutine is a lightweight thread managed by the Go runtime.",
		},
	}

	for i := range assessments {
		err := db.Create(&assessments[i]).Error
		require.NoError(t, err)
	}

	return assessments
}

func TestGetAssessment_Success(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	path := seedTestPath(t, db)
	seedTestAssessment(t, db, path.Milestones[0].ID)

	_, r, _ := setupAssessmentHandler(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/milestones/%d/assessment", path.Milestones[0].ID), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var questions []dto.AssessmentQuestionResponse
	err := json.Unmarshal(w.Body.Bytes(), &questions)
	require.NoError(t, err)
	assert.Equal(t, 3, len(questions))
	assert.NotEmpty(t, questions[0].Question)
	assert.Equal(t, 4, len(questions[0].Options))
}

func TestGetAssessment_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	_, r, _ := setupAssessmentHandler(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/milestones/9999/assessment", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestSubmitAssessment_AllCorrect(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	path := seedTestPath(t, db)
	assessments := seedTestAssessment(t, db, path.Milestones[0].ID)

	_, r, _ := setupAssessmentHandler(db)

	submitReq := dto.SubmitAssessmentRequest{
		Answers: []dto.AnswerSubmission{
			{QuestionID: assessments[0].ID, SelectedIndex: 0},
			{QuestionID: assessments[1].ID, SelectedIndex: 1},
			{QuestionID: assessments[2].ID, SelectedIndex: 1},
		},
	}
	jsonBody, _ := json.Marshal(submitReq)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/milestones/%d/assessment/submit", path.Milestones[0].ID), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result dto.AssessmentResultResponse
	err := json.Unmarshal(w.Body.Bytes(), &result)
	require.NoError(t, err)
	assert.Equal(t, 3, result.Score)
	assert.Equal(t, 3, result.TotalQuestions)
	assert.True(t, result.Passed)
}

func TestSubmitAssessment_PartialCorrect(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	path := seedTestPath(t, db)
	assessments := seedTestAssessment(t, db, path.Milestones[0].ID)

	_, r, _ := setupAssessmentHandler(db)

	submitReq := dto.SubmitAssessmentRequest{
		Answers: []dto.AnswerSubmission{
			{QuestionID: assessments[0].ID, SelectedIndex: 0}, // correct
			{QuestionID: assessments[1].ID, SelectedIndex: 0}, // wrong
			{QuestionID: assessments[2].ID, SelectedIndex: 0}, // wrong
		},
	}
	jsonBody, _ := json.Marshal(submitReq)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/milestones/%d/assessment/submit", path.Milestones[0].ID), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result dto.AssessmentResultResponse
	err := json.Unmarshal(w.Body.Bytes(), &result)
	require.NoError(t, err)
	assert.Equal(t, 1, result.Score)
	assert.Equal(t, 3, result.TotalQuestions)
	assert.False(t, result.Passed)
}
