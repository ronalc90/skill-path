package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/service"
)

// DashboardHandler handles dashboard HTTP requests.
type DashboardHandler struct {
	dashboardService *service.DashboardService
}

func NewDashboardHandler(dashboardService *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService: dashboardService}
}

// GetDashboard godoc
// GET /api/v1/dashboard
func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	userID := c.GetUint("userID")

	resp, err := h.dashboardService.GetDashboard(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to load dashboard"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
