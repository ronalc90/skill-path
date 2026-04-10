package main

import (
	"fmt"
	"log"

	"github.com/ronalc90/skillpath/internal/config"
	"github.com/ronalc90/skillpath/internal/handler"
	"github.com/ronalc90/skillpath/internal/model"
	"github.com/ronalc90/skillpath/internal/repository"
	"github.com/ronalc90/skillpath/internal/router"
	"github.com/ronalc90/skillpath/internal/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := gorm.Open(sqlite.Open("skillpath.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.LearningPath{},
		&model.Milestone{},
		&model.Resource{},
		&model.UserProgress{},
		&model.MilestoneProgress{},
		&model.ResourceProgress{},
		&model.SkillAssessment{},
		&model.AssessmentResult{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repository.Seed(db)

	userRepo := repository.NewUserRepository(db)
	pathRepo := repository.NewPathRepository(db)
	progressRepo := repository.NewProgressRepository(db)
	assessmentRepo := repository.NewAssessmentRepository(db)

	authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiryHrs)
	pathService := service.NewPathService(pathRepo)
	progressService := service.NewProgressService(progressRepo, pathRepo)
	assessmentService := service.NewAssessmentService(assessmentRepo)
	dashboardService := service.NewDashboardService(progressRepo, assessmentRepo)
	searchService := service.NewSearchService(pathRepo)

	authHandler := handler.NewAuthHandler(authService)
	pathHandler := handler.NewPathHandler(pathService)
	progressHandler := handler.NewProgressHandler(progressService)
	assessmentHandler := handler.NewAssessmentHandler(assessmentService)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)
	searchHandler := handler.NewSearchHandler(searchService)

	r := router.Setup(
		cfg,
		authHandler,
		pathHandler,
		progressHandler,
		assessmentHandler,
		dashboardHandler,
		searchHandler,
	)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("SkillPath API starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
