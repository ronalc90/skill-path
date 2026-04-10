package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/service"
)

// AssessmentHandler handles assessment/quiz HTTP requests.
type AssessmentHandler struct {
	assessmentService *service.AssessmentService
}

func NewAssessmentHandler(assessmentService *service.AssessmentService) *AssessmentHandler {
	return &AssessmentHandler{assessmentService: assessmentService}
}

// GetAssessment godoc
// GET /api/v1/milestones/:id/assessment
func (h *AssessmentHandler) GetAssessment(c *gin.Context) {
	milestoneID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid_id", Message: "Invalid milestone ID"})
		return
	}

	questions, err := h.assessmentService.GetAssessment(uint(milestoneID))
	if err != nil {
		if errors.Is(err, service.ErrNoAssessment) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_found", Message: "No assessment found for this milestone"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to get assessment"})
		return
	}

	c.JSON(http.StatusOK, questions)
}

// SubmitAssessment godoc
// POST /api/v1/milestones/:id/assessment/submit
func (h *AssessmentHandler) SubmitAssessment(c *gin.Context) {
	userID := c.GetUint("userID")
	milestoneID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid_id", Message: "Invalid milestone ID"})
		return
	}

	var req dto.SubmitAssessmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "validation_error", Message: err.Error()})
		return
	}

	result, err := h.assessmentService.SubmitAssessment(userID, uint(milestoneID), req)
	if err != nil {
		if errors.Is(err, service.ErrNoAssessment) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_found", Message: "No assessment found for this milestone"})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to submit assessment"})
		return
	}

	c.JSON(http.StatusOK, result)
}
