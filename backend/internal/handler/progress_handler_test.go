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

func setupProgressHandler(db *gorm.DB) (*handler.ProgressHandler, *gin.Engine) {
	gin.SetMode(gin.TestMode)

	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)
	progressHandler := handler.NewProgressHandler(progressService)

	r := gin.New()

	// Middleware to inject userID for testing
	authMiddleware := func(c *gin.Context) {
		c.Set("userID", uint(1))
		c.Next()
	}

	r.POST("/api/v1/paths/:slug/start", authMiddleware, progressHandler.StartPath)
	r.GET("/api/v1/paths/:slug/progress", authMiddleware, progressHandler.GetPathProgress)
	r.PUT("/api/v1/paths/:slug/milestones/:milestoneId/complete", authMiddleware, progressHandler.CompleteMilestone)
	r.GET("/api/v1/my-paths", authMiddleware, progressHandler.GetMyPaths)

	return progressHandler, r
}

func seedTestUser(t *testing.T, db *gorm.DB) *model.User {
	t.Helper()
	hashed, _ := hash.HashPassword("password123")
	user := &model.User{
		Email:        "progress-test@example.com",
		PasswordHash: hashed,
		DisplayName:  "Progress Test User",
	}
	err := db.Create(user).Error
	require.NoError(t, err)
	return user
}

func TestStartPath_Success(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	user := seedTestUser(t, db)
	path := seedTestPath(t, db)

	_, r := setupProgressHandler(db)

	// Override the middleware to use the real user ID
	r.Handlers = nil
	authMiddleware := func(c *gin.Context) {
		c.Set("userID", user.ID)
		c.Next()
	}

	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)
	progressHandler := handler.NewProgressHandler(progressService)

	r2 := gin.New()
	r2.POST("/api/v1/paths/:slug/start", authMiddleware, progressHandler.StartPath)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/paths/%s/start", path.Slug), nil)
	r2.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp dto.UserPathProgress
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, path.ID, resp.PathID)
	assert.Equal(t, float64(0), resp.ProgressPct)
}

func TestStartPath_AlreadyEnrolled(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	user := seedTestUser(t, db)
	path := seedTestPath(t, db)

	authMiddleware := func(c *gin.Context) {
		c.Set("userID", user.ID)
		c.Next()
	}

	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)
	progressHandler := handler.NewProgressHandler(progressService)

	r := gin.New()
	r.POST("/api/v1/paths/:slug/start", authMiddleware, progressHandler.StartPath)

	// First enrollment
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/paths/%s/start", path.Slug), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Duplicate enrollment
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", fmt.Sprintf("/api/v1/paths/%s/start", path.Slug), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestStartPath_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	user := seedTestUser(t, db)

	authMiddleware := func(c *gin.Context) {
		c.Set("userID", user.ID)
		c.Next()
	}

	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)
	progressHandler := handler.NewProgressHandler(progressService)

	r := gin.New()
	r.POST("/api/v1/paths/:slug/start", authMiddleware, progressHandler.StartPath)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/paths/nonexistent-path/start", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCompleteMilestone_Success(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	user := seedTestUser(t, db)
	path := seedTestPath(t, db)

	authMiddleware := func(c *gin.Context) {
		c.Set("userID", user.ID)
		c.Next()
	}

	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepo, pathRepo)
	progressHandler := handler.NewProgressHandler(progressService)

	r := gin.New()
	r.POST("/api/v1/paths/:slug/start", authMiddleware, progressHandler.StartPath)
	r.PUT("/api/v1/paths/:slug/milestones/:milestoneId/complete", authMiddleware, progressHandler.CompleteMilestone)
	r.GET("/api/v1/paths/:slug/progress", authMiddleware, progressHandler.GetPathProgress)

	// Enroll first
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/paths/%s/start", path.Slug), nil)
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// Complete first milestone
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/v1/paths/%s/milestones/%d/complete", path.Slug, path.Milestones[0].ID), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Check progress
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/paths/%s/progress", path.Slug), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var progress dto.UserPathProgress
	err := json.Unmarshal(w.Body.Bytes(), &progress)
	require.NoError(t, err)
	assert.Equal(t, float64(50), progress.ProgressPct)
}
