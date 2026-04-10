package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/service"
)

// resolvePathID resolves a slug parameter to a numeric path ID.
func resolvePathID(c *gin.Context, progressService *service.ProgressService) (uint, bool) {
	slug := c.Param("slug")
	pathID, err := progressService.FindPathIDBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_found", Message: "Learning path not found"})
		return 0, false
	}
	return pathID, true
}

// ProgressHandler handles progress tracking HTTP requests.
type ProgressHandler struct {
	progressService *service.ProgressService
}

func NewProgressHandler(progressService *service.ProgressService) *ProgressHandler {
	return &ProgressHandler{progressService: progressService}
}

// StartPath godoc
// POST /api/v1/paths/:slug/start
func (h *ProgressHandler) StartPath(c *gin.Context) {
	userID := c.GetUint("userID")
	pathID, ok := resolvePathID(c, h.progressService)
	if !ok {
		return
	}

	resp, err := h.progressService.StartPath(userID, pathID)
	if err != nil {
		if errors.Is(err, service.ErrAlreadyEnrolled) {
			c.JSON(http.StatusConflict, dto.ErrorResponse{Error: "already_enrolled", Message: "Already enrolled in this path"})
			return
		}
		if errors.Is(err, service.ErrPathNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_found", Message: "Learning path not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to start path"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetPathProgress godoc
// GET /api/v1/paths/:slug/progress
func (h *ProgressHandler) GetPathProgress(c *gin.Context) {
	userID := c.GetUint("userID")
	pathID, ok := resolvePathID(c, h.progressService)
	if !ok {
		return
	}

	resp, err := h.progressService.GetPathProgress(userID, pathID)
	if err != nil {
		if errors.Is(err, service.ErrNotEnrolled) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_enrolled", Message: "Not enrolled in this path"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to get progress"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CompleteMilestone godoc
// PUT /api/v1/paths/:slug/milestones/:milestoneId/complete
func (h *ProgressHandler) CompleteMilestone(c *gin.Context) {
	userID := c.GetUint("userID")
	pathID, ok := resolvePathID(c, h.progressService)
	if !ok {
		return
	}

	milestoneID, err := strconv.ParseUint(c.Param("milestoneId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid_id", Message: "Invalid milestone ID"})
		return
	}

	if err := h.progressService.CompleteMilestone(userID, pathID, uint(milestoneID)); err != nil {
		if errors.Is(err, service.ErrNotEnrolled) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_enrolled", Message: "Not enrolled in this path"})
			return
		}
		if errors.Is(err, service.ErrResourcesNotCompleted) {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "resources_incomplete", Message: "Debes completar todos los recursos del milestone antes de marcarlo como completado"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to complete milestone"})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Milestone completed successfully"})
}

// CompleteResource godoc
// PUT /api/v1/paths/:slug/resources/:resourceId/complete
func (h *ProgressHandler) CompleteResource(c *gin.Context) {
	userID := c.GetUint("userID")
	pathID, ok := resolvePathID(c, h.progressService)
	if !ok {
		return
	}

	resourceID, err := strconv.ParseUint(c.Param("resourceId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid_id", Message: "Invalid resource ID"})
		return
	}

	var req dto.CompleteResourceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		req.Rating = 0
	}

	if err := h.progressService.CompleteResource(userID, pathID, uint(resourceID), req.Rating); err != nil {
		if errors.Is(err, service.ErrNotEnrolled) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_enrolled", Message: "Not enrolled in this path"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to complete resource"})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Resource completed successfully"})
}

// GetMyPaths godoc
// GET /api/v1/my-paths
func (h *ProgressHandler) GetMyPaths(c *gin.Context) {
	userID := c.GetUint("userID")

	paths, err := h.progressService.GetMyPaths(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to get enrolled paths"})
		return
	}

	c.JSON(http.StatusOK, paths)
}
