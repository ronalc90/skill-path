package handler_test

import (
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
	"github.com/ronalc90/skillpath/pkg/hash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupDashboardHandler(db *gorm.DB, userID uint) (*handler.DashboardHandler, *gin.Engine) {
	gin.SetMode(gin.TestMode)

	progressRepo := repository.NewProgressRepository(db)
	assessmentRepo := repository.NewAssessmentRepository(db)
	dashboardService := service.NewDashboardService(progressRepo, assessmentRepo)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	authMiddleware := func(c *gin.Context) {
		c.Set("userID", userID)
		c.Next()
	}

	r := gin.New()
	r.GET("/api/v1/dashboard", authMiddleware, dashboardHandler.GetDashboard)

	return dashboardHandler, r
}

func TestGetDashboard_Empty(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	hashed, _ := hash.HashPassword("password123")
	user := &model.User{
		Email:        "dashboard-test@example.com",
		PasswordHash: hashed,
		DisplayName:  "Dashboard User",
	}
	db.Create(user)

	_, r := setupDashboardHandler(db, user.ID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/dashboard", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.DashboardResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, resp.PathsInProgress)
	assert.Equal(t, 0, resp.PathsCompleted)
	assert.Equal(t, float64(0), resp.TotalHoursLearned)
}

func TestGetDashboard_WithProgress(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	hashed, _ := hash.HashPassword("password123")
	user := &model.User{
		Email:        "dashboard-progress@example.com",
		PasswordHash: hashed,
		DisplayName:  "Dashboard Progress User",
	}
	db.Create(user)

	path := seedTestPath(t, db)

	// Enroll user in path
	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)

	_, err := progressService.StartPath(user.ID, path.ID)
	require.NoError(t, err)

	// Complete a milestone
	err = progressService.CompleteMilestone(user.ID, path.ID, path.Milestones[0].ID)
	require.NoError(t, err)

	_, r := setupDashboardHandler(db, user.ID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/dashboard", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.DashboardResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, 1, resp.PathsInProgress)
	assert.Equal(t, 0, resp.PathsCompleted)
	assert.Greater(t, resp.TotalHoursLearned, float64(0))
	assert.NotEmpty(t, resp.EnrolledPaths)

	// Verify enrolled path progress
	enrolledPath := resp.EnrolledPaths[0]
	assert.Equal(t, path.ID, enrolledPath.PathID)
	assert.Equal(t, float64(50), enrolledPath.ProgressPct)

	_ = fmt.Sprintf("Path progress: %.0f%%", enrolledPath.ProgressPct)
}
