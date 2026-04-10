package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/service"
)

// SearchHandler handles search HTTP requests.
type SearchHandler struct {
	searchService *service.SearchService
}

func NewSearchHandler(searchService *service.SearchService) *SearchHandler {
	return &SearchHandler{searchService: searchService}
}

// Search godoc
// GET /api/v1/search
func (h *SearchHandler) Search(c *gin.Context) {
	query := c.Query("q")

	resp, err := h.searchService.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "internal_error", Message: "Search failed"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
