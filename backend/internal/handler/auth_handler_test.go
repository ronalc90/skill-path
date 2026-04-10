package handler_test

import (
	"bytes"
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
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Skipf("Skipping test: database not available: %v", err)
	}

	// Auto-migrate test tables
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

func cleanupTestDB(t *testing.T, db *gorm.DB) {
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

func setupAuthHandler(db *gorm.DB) (*handler.AuthHandler, *gin.Engine) {
	gin.SetMode(gin.TestMode)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, "test-secret", 72)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.New()
	r.POST("/api/v1/auth/register", authHandler.Register)
	r.POST("/api/v1/auth/login", authHandler.Login)

	return authHandler, r
}

func TestRegister_Success(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	_, r := setupAuthHandler(db)

	body := dto.RegisterRequest{
		Email:       "test@example.com",
		Password:    "password123",
		DisplayName: "Test User",
	}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp dto.AuthResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Token)
	assert.Equal(t, "test@example.com", resp.User.Email)
	assert.Equal(t, "Test User", resp.User.DisplayName)
}

func TestRegister_DuplicateEmail(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	_, r := setupAuthHandler(db)

	body := dto.RegisterRequest{
		Email:       "duplicate@example.com",
		Password:    "password123",
		DisplayName: "First User",
	}
	jsonBody, _ := json.Marshal(body)

	// First registration
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Duplicate registration
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestRegister_InvalidEmail(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	_, r := setupAuthHandler(db)

	body := map[string]string{
		"email":        "not-an-email",
		"password":     "password123",
		"display_name": "Test",
	}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin_Success(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	_, r := setupAuthHandler(db)

	// Register first
	regBody := dto.RegisterRequest{
		Email:       "login@example.com",
		Password:    "password123",
		DisplayName: "Login User",
	}
	jsonBody, _ := json.Marshal(regBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code)

	// Login
	loginBody := dto.LoginRequest{
		Email:    "login@example.com",
		Password: "password123",
	}
	jsonBody, _ = json.Marshal(loginBody)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp dto.AuthResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Token)
	assert.Equal(t, "login@example.com", resp.User.Email)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	db := setupTestDB(t)
	defer cleanupTestDB(t, db)

	_, r := setupAuthHandler(db)

	loginBody := dto.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "wrongpassword",
	}
	jsonBody, _ := json.Marshal(loginBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
