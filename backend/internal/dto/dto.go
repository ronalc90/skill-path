package dto

import "time"

// --- Auth DTOs ---

// RegisterRequest represents the registration request body.
type RegisterRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	DisplayName string `json:"display_name" binding:"required,min=2,max=100"`
}

// LoginRequest represents the login request body.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse is returned after successful authentication.
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// UserResponse represents a sanitized user in responses.
type UserResponse struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	Bio         string `json:"bio,omitempty"`
	GithubURL   string `json:"github_url,omitempty"`
	LinkedinURL string `json:"linkedin_url,omitempty"`
}

// UpdateProfileRequest is used to update user profile fields.
type UpdateProfileRequest struct {
	DisplayName string `json:"display_name" binding:"omitempty,min=2,max=100"`
	AvatarURL   string `json:"avatar_url" binding:"omitempty,url"`
	Bio         string `json:"bio" binding:"omitempty,max=500"`
	GithubURL   string `json:"github_url" binding:"omitempty,url"`
	LinkedinURL string `json:"linkedin_url" binding:"omitempty,url"`
}

// --- Path DTOs ---

// PathListResponse is a paginated list of learning paths.
type PathListResponse struct {
	Paths      []PathSummary `json:"paths"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	PageSize   int           `json:"page_size"`
	TotalPages int           `json:"total_pages"`
}

// PathSummary is a condensed representation of a learning path.
type PathSummary struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	Difficulty     string `json:"difficulty"`
	EstimatedHours int    `json:"estimated_hours"`
	Category       string `json:"category"`
	IconURL        string `json:"icon_url,omitempty"`
	MilestoneCount int    `json:"milestone_count"`
}

// PathDetailResponse is the full detail of a learning path with milestones and resources.
type PathDetailResponse struct {
	ID             uint                `json:"id"`
	Title          string              `json:"title"`
	Slug           string              `json:"slug"`
	Description    string              `json:"description"`
	Difficulty     string              `json:"difficulty"`
	EstimatedHours int                 `json:"estimated_hours"`
	Category       string              `json:"category"`
	IconURL        string              `json:"icon_url,omitempty"`
	IsPublished    bool                `json:"is_published"`
	Milestones     []MilestoneResponse `json:"milestones"`
}

// MilestoneResponse represents a milestone with its resources.
type MilestoneResponse struct {
	ID             uint               `json:"id"`
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	OrderIndex     int                `json:"order_index"`
	EstimatedHours float64            `json:"estimated_hours"`
	Resources      []ResourceResponse `json:"resources"`
}

// ResourceResponse represents a single learning resource.
type ResourceResponse struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	Type             string `json:"type"`
	Provider         string `json:"provider,omitempty"`
	IsFree           bool   `json:"is_free"`
	EstimatedMinutes int    `json:"estimated_minutes"`
	OrderIndex       int    `json:"order_index"`
}

// CategoryResponse represents a category with its path count.
type CategoryResponse struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// --- Progress DTOs ---

// UserPathProgress represents the progress of a user on a path.
type UserPathProgress struct {
	PathID         uint                      `json:"path_id"`
	Path           PathSummary               `json:"path"`
	StartedAt      time.Time                 `json:"started_at"`
	CompletedAt    *time.Time                `json:"completed_at,omitempty"`
	LastActivityAt time.Time                 `json:"last_activity_at"`
	ProgressPct    float64                   `json:"progress_pct"`
	Milestones     []MilestoneProgressDetail `json:"milestones,omitempty"`
}

// MilestoneProgressDetail represents a single milestone's progress.
type MilestoneProgressDetail struct {
	MilestoneID uint   `json:"milestone_id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	OrderIndex  int    `json:"order_index"`
}

// CompleteMilestoneRequest is sent to mark a milestone as complete.
type CompleteMilestoneRequest struct {
	MilestoneID uint `json:"milestone_id" binding:"required"`
}

// CompleteResourceRequest is sent to mark a resource as complete.
type CompleteResourceRequest struct {
	ResourceID uint `json:"resource_id" binding:"required"`
	Rating     int  `json:"rating" binding:"omitempty,min=1,max=5"`
}

// --- Assessment DTOs ---

// AssessmentQuestionResponse is a single quiz question (without the correct answer).
type AssessmentQuestionResponse struct {
	ID       uint     `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

// SubmitAssessmentRequest is the request body for submitting quiz answers.
type SubmitAssessmentRequest struct {
	Answers []AnswerSubmission `json:"answers" binding:"required,dive"`
}

// AnswerSubmission represents a user's answer to one question.
type AnswerSubmission struct {
	QuestionID    uint `json:"question_id" binding:"required"`
	SelectedIndex int  `json:"selected_index" binding:"min=0"`
}

// AssessmentResultResponse is the result of a quiz attempt.
type AssessmentResultResponse struct {
	Score          int              `json:"score"`
	TotalQuestions int              `json:"total_questions"`
	Passed         bool             `json:"passed"`
	Details        []AnswerDetail   `json:"details"`
}

// AnswerDetail shows whether a specific answer was correct.
type AnswerDetail struct {
	QuestionID    uint   `json:"question_id"`
	IsCorrect     bool   `json:"is_correct"`
	CorrectAnswer int    `json:"correct_answer"`
	Explanation   string `json:"explanation"`
}

// --- Dashboard DTOs ---

// DashboardResponse aggregates the user's learning dashboard data.
type DashboardResponse struct {
	PathsInProgress   int                `json:"paths_in_progress"`
	PathsCompleted    int                `json:"paths_completed"`
	TotalHoursLearned float64            `json:"total_hours_learned"`
	CompletionRate    float64            `json:"completion_rate"`
	RecentActivity    []RecentActivity   `json:"recent_activity"`
	SkillRadar        []SkillRadarPoint  `json:"skill_radar"`
	EnrolledPaths     []UserPathProgress `json:"enrolled_paths"`
}

// RecentActivity represents a single recent learning activity.
type RecentActivity struct {
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	PathTitle string    `json:"path_title"`
	Timestamp time.Time `json:"timestamp"`
}

// SkillRadarPoint represents a data point for the skill radar chart.
type SkillRadarPoint struct {
	Category string  `json:"category"`
	Score    float64 `json:"score"`
	MaxScore float64 `json:"max_score"`
}

// --- Search DTOs ---

// SearchResponse is the combined search result.
type SearchResponse struct {
	Paths     []PathSummary      `json:"paths"`
	Resources []ResourceResponse `json:"resources"`
}

// ErrorResponse is a standardized error response.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// MessageResponse is a simple message response.
type MessageResponse struct {
	Message string `json:"message"`
}
