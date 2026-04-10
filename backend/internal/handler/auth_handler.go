package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/service"
)

// AuthHandler handles authentication HTTP requests.
type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// POST /api/v1/auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "validation_error", Message: err.Error()})
		return
	}

	resp, err := h.authService.Register(req)
	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			c.JSON(http.StatusConflict, dto.ErrorResponse{Error: "email_exists", Message: "Email already registered"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Login godoc
// POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "validation_error", Message: err.Error()})
		return
	}

	resp, err := h.authService.Login(req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "invalid_credentials", Message: "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Login failed"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetMe godoc
// GET /api/v1/auth/me
func (h *AuthHandler) GetMe(c *gin.Context) {
	userID := c.GetUint("userID")

	resp, err := h.authService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_found", Message: "User not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateProfile godoc
// PUT /api/v1/auth/profile
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "validation_error", Message: err.Error()})
		return
	}

	resp, err := h.authService.UpdateProfile(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
