package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Difficulty represents the difficulty level of a learning path.
type Difficulty string

const (
	DifficultyBeginner     Difficulty = "beginner"
	DifficultyIntermediate Difficulty = "intermediate"
	DifficultyAdvanced     Difficulty = "advanced"
)

// ResourceType represents the type of learning resource.
type ResourceType string

const (
	ResourceTypeVideo    ResourceType = "video"
	ResourceTypeArticle  ResourceType = "article"
	ResourceTypeCourse   ResourceType = "course"
	ResourceTypeBook     ResourceType = "book"
	ResourceTypeTutorial ResourceType = "tutorial"
	ResourceTypeExercise ResourceType = "exercise"
)

// MilestoneStatus represents the progress status of a milestone.
type MilestoneStatus string

const (
	StatusNotStarted MilestoneStatus = "not_started"
	StatusInProgress MilestoneStatus = "in_progress"
	StatusCompleted  MilestoneStatus = "completed"
)

// JSONSlice is a custom type for storing JSON arrays in PostgreSQL.
type JSONSlice []string

func (j JSONSlice) Value() (driver.Value, error) {
	if j == nil {
		return "[]", nil
	}
	return json.Marshal(j)
}

func (j *JSONSlice) Scan(value interface{}) error {
	if value == nil {
		*j = JSONSlice{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("failed to scan JSONSlice: invalid type")
	}
	return json.Unmarshal(bytes, j)
}

// User represents a registered user.
type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Email        string         `gorm:"uniqueIndex;size:255;not null" json:"email"`
	PasswordHash string         `gorm:"size:255;not null" json:"-"`
	DisplayName  string         `gorm:"size:100;not null" json:"display_name"`
	AvatarURL    string         `gorm:"size:500" json:"avatar_url,omitempty"`
	Bio          string         `gorm:"size:500" json:"bio,omitempty"`
	GithubURL    string         `gorm:"size:255" json:"github_url,omitempty"`
	LinkedinURL  string         `gorm:"size:255" json:"linkedin_url,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// LearningPath represents a complete learning path.
type LearningPath struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Title          string     `gorm:"size:200;not null" json:"title"`
	Description    string     `gorm:"type:text;not null" json:"description"`
	Slug           string     `gorm:"uniqueIndex;size:200;not null" json:"slug"`
	Difficulty     Difficulty `gorm:"size:20;not null;default:'beginner'" json:"difficulty"`
	EstimatedHours int        `gorm:"not null;default:0" json:"estimated_hours"`
	Category       string     `gorm:"size:100;not null" json:"category"`
	IconURL        string     `gorm:"size:500" json:"icon_url,omitempty"`
	IsPublished    bool       `gorm:"not null;default:false" json:"is_published"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	Milestones     []Milestone `gorm:"foreignKey:PathID" json:"milestones,omitempty"`
}

// Milestone represents a milestone within a learning path.
type Milestone struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	PathID         uint       `gorm:"not null;index" json:"path_id"`
	Title          string     `gorm:"size:200;not null" json:"title"`
	Description    string     `gorm:"type:text" json:"description"`
	OrderIndex     int        `gorm:"not null;default:0" json:"order_index"`
	EstimatedHours float64    `gorm:"not null;default:0" json:"estimated_hours"`
	Resources      []Resource `gorm:"foreignKey:MilestoneID" json:"resources,omitempty"`
}

// Resource represents a learning resource within a milestone.
type Resource struct {
	ID               uint         `gorm:"primaryKey" json:"id"`
	MilestoneID      uint         `gorm:"not null;index" json:"milestone_id"`
	Title            string       `gorm:"size:300;not null" json:"title"`
	URL              string       `gorm:"size:500;not null" json:"url"`
	Type             ResourceType `gorm:"size:20;not null" json:"type"`
	Provider         string       `gorm:"size:100" json:"provider,omitempty"`
	IsFree           bool         `gorm:"not null;default:true" json:"is_free"`
	EstimatedMinutes int          `gorm:"not null;default:0" json:"estimated_minutes"`
	OrderIndex       int          `gorm:"not null;default:0" json:"order_index"`
}

// UserProgress tracks a user's enrollment and progress in a learning path.
type UserProgress struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	UserID         uint       `gorm:"not null;index;uniqueIndex:idx_user_path" json:"user_id"`
	PathID         uint       `gorm:"not null;index;uniqueIndex:idx_user_path" json:"path_id"`
	StartedAt      time.Time  `gorm:"not null" json:"started_at"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
	LastActivityAt time.Time  `gorm:"not null" json:"last_activity_at"`
	User           User         `gorm:"foreignKey:UserID" json:"-"`
	Path           LearningPath `gorm:"foreignKey:PathID" json:"path,omitempty"`
	MilestoneProgresses []MilestoneProgress `gorm:"foreignKey:UserProgressID" json:"milestone_progresses,omitempty"`
	ResourceProgresses  []ResourceProgress  `gorm:"foreignKey:UserProgressID" json:"resource_progresses,omitempty"`
}

// MilestoneProgress tracks a user's progress on a specific milestone.
type MilestoneProgress struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	UserProgressID uint            `gorm:"not null;index;uniqueIndex:idx_up_milestone" json:"user_progress_id"`
	MilestoneID    uint            `gorm:"not null;index;uniqueIndex:idx_up_milestone" json:"milestone_id"`
	Status         MilestoneStatus `gorm:"size:20;not null;default:'not_started'" json:"status"`
	CompletedAt    *time.Time      `json:"completed_at,omitempty"`
	Milestone      Milestone       `gorm:"foreignKey:MilestoneID" json:"milestone,omitempty"`
}

// ResourceProgress tracks a user's completion of a specific resource.
type ResourceProgress struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	UserProgressID uint       `gorm:"not null;index;uniqueIndex:idx_up_resource" json:"user_progress_id"`
	ResourceID     uint       `gorm:"not null;index;uniqueIndex:idx_up_resource" json:"resource_id"`
	IsCompleted    bool       `gorm:"not null;default:false" json:"is_completed"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
	Rating         int        `gorm:"default:0" json:"rating"`
	Resource       Resource   `gorm:"foreignKey:ResourceID" json:"resource,omitempty"`
}

// SkillAssessment represents a quiz question for a milestone.
type SkillAssessment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	MilestoneID   uint      `gorm:"not null;index" json:"milestone_id"`
	Question      string    `gorm:"type:text;not null" json:"question"`
	Options       JSONSlice `gorm:"type:text;not null" json:"options"`
	CorrectAnswer int       `gorm:"not null" json:"correct_answer"`
	Explanation   string    `gorm:"type:text" json:"explanation"`
}

// AssessmentResult stores a user's quiz attempt results.
type AssessmentResult struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	UserID         uint       `gorm:"not null;index" json:"user_id"`
	MilestoneID    uint       `gorm:"not null;index" json:"milestone_id"`
	Score          int        `gorm:"not null" json:"score"`
	TotalQuestions int        `gorm:"not null" json:"total_questions"`
	PassedAt       *time.Time `json:"passed_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}
