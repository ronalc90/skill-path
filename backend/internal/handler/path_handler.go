package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/service"
)

// PathHandler handles learning path HTTP requests.
type PathHandler struct {
	pathService *service.PathService
}

func NewPathHandler(pathService *service.PathService) *PathHandler {
	return &PathHandler{pathService: pathService}
}

// ListPaths godoc
// GET /api/v1/paths
func (h *PathHandler) ListPaths(c *gin.Context) {
	category := c.Query("category")
	difficulty := c.Query("difficulty")
	search := c.Query("search")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))

	resp, err := h.pathService.ListPaths(category, difficulty, search, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to list paths"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetPathBySlug godoc
// GET /api/v1/paths/:slug
func (h *PathHandler) GetPathBySlug(c *gin.Context) {
	slug := c.Param("slug")

	resp, err := h.pathService.GetPathBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not_found", Message: "Learning path not found"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetCategories godoc
// GET /api/v1/paths/categories
func (h *PathHandler) GetCategories(c *gin.Context) {
	categories, err := h.pathService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Failed to get categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}
