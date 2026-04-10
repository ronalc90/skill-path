package router

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ronalc90/skillpath/internal/config"
	"github.com/ronalc90/skillpath/internal/handler"
	"github.com/ronalc90/skillpath/internal/middleware"
)

// Setup configures all routes and middleware for the Gin engine.
func Setup(
	cfg *config.Config,
	authHandler *handler.AuthHandler,
	pathHandler *handler.PathHandler,
	progressHandler *handler.ProgressHandler,
	assessmentHandler *handler.AssessmentHandler,
	dashboardHandler *handler.DashboardHandler,
	searchHandler *handler.SearchHandler,
) *gin.Engine {
	gin.SetMode(cfg.GinMode)

	r := gin.New()

	// Global middleware
	r.Use(middleware.Recovery())
	r.Use(middleware.Logging())
	r.Use(middleware.Metrics())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(cfg.CORSOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	api := r.Group("/api/v1")

	// Auth routes (public)
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Auth routes (protected)
	authProtected := api.Group("/auth")
	authProtected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		authProtected.GET("/me", authHandler.GetMe)
		authProtected.PUT("/profile", authHandler.UpdateProfile)
	}

	// Path routes (public)
	paths := api.Group("/paths")
	{
		paths.GET("", pathHandler.ListPaths)
		paths.GET("/categories", pathHandler.GetCategories)
		paths.GET("/:slug", pathHandler.GetPathBySlug)
	}

	// Progress routes (protected) - use same :slug param to avoid Gin conflict
	protectedPaths := api.Group("/paths")
	protectedPaths.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		protectedPaths.POST("/:slug/start", progressHandler.StartPath)
		protectedPaths.GET("/:slug/progress", progressHandler.GetPathProgress)
		protectedPaths.PUT("/:slug/milestones/:milestoneId/complete", progressHandler.CompleteMilestone)
		protectedPaths.PUT("/:slug/resources/:resourceId/complete", progressHandler.CompleteResource)
	}

	// My paths (protected)
	api.GET("/my-paths", middleware.AuthMiddleware(cfg.JWTSecret), progressHandler.GetMyPaths)

	// Assessment routes (protected)
	milestones := api.Group("/milestones")
	milestones.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		milestones.GET("/:id/assessment", assessmentHandler.GetAssessment)
		milestones.POST("/:id/assessment/submit", assessmentHandler.SubmitAssessment)
	}

	// Dashboard (protected)
	api.GET("/dashboard", middleware.AuthMiddleware(cfg.JWTSecret), dashboardHandler.GetDashboard)

	// Search (public)
	api.GET("/search", searchHandler.Search)

	return r
}
