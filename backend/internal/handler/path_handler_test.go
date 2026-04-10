package handler_test

import (
	"encoding/json"
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

func setupPathHandler(db *gorm.DB) (*handler.PathHandler, *gin.Engine) {
	gin.SetMode(gin.TestMode)

	pathRepo := repository.NewPathRepository(db)
	pathService := service.NewPathService(pathRepo)
	pathHandler := handler.NewPathHandler(pathService)

	r := gin.New()
	r.GET("/api/v1/paths", pathHandler.ListPaths)
	r.GET("/api/v1/paths/categories", pathHandler.GetCategories)
	r.GET("/api/v1/paths/:slug", pathHandler.GetPathBySlug)

	return pathHandler, r
}

func seedTestPath(t *testing.T, db *gorm.DB) *model.LearningPath {
	t.Helper()
	path := &model.LearningPath{
		Title:          "Test Path",
		Description:    "A test learning path",
		Slug:           "test-path",
		Difficulty:     model.DifficultyBeginner,
		EstimatedHours: 10,
		Category:       "Testing",
		IsPublished:    true,
		Milestones: []model.Milestone{
			{
				Title:          "Milestone 1",
				Description:    "First milestone",
				OrderIndex:     1,
				EstimatedHours: 5,
				Resources: []model.Resource{
					{Title: "Resource 1", URL: "https://example.com/r1", Type: model.ResourceTypeArticle, IsFree: true, EstimatedMinutes: 30, OrderIndex: 1},
				},
			},
			{
				Title:          "Milestone 2",
				Description:    "Second milestone",
				OrderIndex:     2,
				EstimatedHours: 5,
				Resources: []model.Resource{
					{Title: "Resource 2", URL: "https://example.com/r2", Type: model.ResourceTypeVideo, IsFree: true, EstimatedMinutes: 60, OrderIndex: 1},
				},
			},
		},
	}

	err := db.Create(path).Error
	require.NoError(t, err)
	return path
}

func TestListPaths_Success(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	seedTestPath(t, db)
	_, r := setupPathHandler(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/paths", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.PathListResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(resp.Paths), 1)
	assert.Equal(t, "Test Path", resp.Paths[0].Title)
}

func TestListPaths_FilterByCategory(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	seedTestPath(t, db)
	_, r := setupPathHandler(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/paths?category=Testing", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.PathListResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, 1, len(resp.Paths))
	assert.Equal(t, "Testing", resp.Paths[0].Category)
}

func TestListPaths_FilterByNonExistentCategory(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	seedTestPath(t, db)
	_, r := setupPathHandler(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/paths?category=NonExistent", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.PathListResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, 0, len(resp.Paths))
}

func TestGetPathBySlug_Success(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	seedTestPath(t, db)
	_, r := setupPathHandler(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/paths/test-path", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.PathDetailResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, "Test Path", resp.Title)
	assert.Equal(t, 2, len(resp.Milestones))
	assert.Equal(t, 1, len(resp.Milestones[0].Resources))
}

func TestGetPathBySlug_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	_, r := setupPathHandler(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/paths/nonexistent-path", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetCategories_Success(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	seedTestPath(t, db)
	_, r := setupPathHandler(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/paths/categories", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var categories []dto.CategoryResponse
	err := json.Unmarshal(w.Body.Bytes(), &categories)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(categories), 1)
}
