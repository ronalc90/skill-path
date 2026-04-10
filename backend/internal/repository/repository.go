package repository

import (
	"github.com/ronalc90/skillpath/internal/model"
	"gorm.io/gorm"
)

// UserRepository handles database operations for users.
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// PathRepository handles database operations for learning paths.
type PathRepository struct {
	db *gorm.DB
}

func NewPathRepository(db *gorm.DB) *PathRepository {
	return &PathRepository{db: db}
}

func (r *PathRepository) FindAll(category, difficulty, search string, page, pageSize int) ([]model.LearningPath, int64, error) {
	var paths []model.LearningPath
	var total int64

	query := r.db.Model(&model.LearningPath{}).Where("is_published = ?", true)

	if category != "" {
		query = query.Where("category = ?", category)
	}
	if difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}
	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("title LIKE ? OR description LIKE ?", searchTerm, searchTerm)
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("Milestones").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&paths).Error

	return paths, total, err
}

func (r *PathRepository) FindBySlug(slug string) (*model.LearningPath, error) {
	var path model.LearningPath
	err := r.db.Where("slug = ? AND is_published = ?", slug, true).
		Preload("Milestones", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_index ASC")
		}).
		Preload("Milestones.Resources", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_index ASC")
		}).
		First(&path).Error
	if err != nil {
		return nil, err
	}
	return &path, nil
}

func (r *PathRepository) FindByID(id uint) (*model.LearningPath, error) {
	var path model.LearningPath
	err := r.db.Preload("Milestones").First(&path, id).Error
	if err != nil {
		return nil, err
	}
	return &path, nil
}

func (r *PathRepository) GetCategories() ([]struct {
	Category string
	Count    int64
}, error) {
	var results []struct {
		Category string
		Count    int64
	}
	err := r.db.Model(&model.LearningPath{}).
		Select("category, count(*) as count").
		Where("is_published = ?", true).
		Group("category").
		Order("count DESC").
		Find(&results).Error
	return results, err
}

func (r *PathRepository) SearchPaths(query string) ([]model.LearningPath, error) {
	var paths []model.LearningPath
	searchTerm := "%" + query + "%"
	err := r.db.Where("is_published = ? AND (title LIKE ? OR description LIKE ?)", true, searchTerm, searchTerm).
		Preload("Milestones").
		Limit(20).
		Find(&paths).Error
	return paths, err
}

func (r *PathRepository) SearchResources(query string) ([]model.Resource, error) {
	var resources []model.Resource
	searchTerm := "%" + query + "%"
	err := r.db.Where("title LIKE ?", searchTerm).
		Limit(20).
		Find(&resources).Error
	return resources, err
}

// ProgressRepository handles database operations for user progress.
type ProgressRepository struct {
	db *gorm.DB
}

func NewProgressRepository(db *gorm.DB) *ProgressRepository {
	return &ProgressRepository{db: db}
}

func (r *ProgressRepository) FindUserProgress(userID, pathID uint) (*model.UserProgress, error) {
	var progress model.UserProgress
	err := r.db.Where("user_id = ? AND path_id = ?", userID, pathID).
		Preload("MilestoneProgresses").
		Preload("MilestoneProgresses.Milestone").
		Preload("ResourceProgresses").
		First(&progress).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

func (r *ProgressRepository) CreateUserProgress(progress *model.UserProgress) error {
	return r.db.Create(progress).Error
}

func (r *ProgressRepository) UpdateUserProgress(progress *model.UserProgress) error {
	return r.db.Save(progress).Error
}

func (r *ProgressRepository) FindAllUserPaths(userID uint) ([]model.UserProgress, error) {
	var progresses []model.UserProgress
	err := r.db.Where("user_id = ?", userID).
		Preload("Path").
		Preload("Path.Milestones").
		Preload("MilestoneProgresses").
		Preload("MilestoneProgresses.Milestone").
		Order("last_activity_at DESC").
		Find(&progresses).Error
	return progresses, err
}

func (r *ProgressRepository) FindMilestoneProgress(userProgressID, milestoneID uint) (*model.MilestoneProgress, error) {
	var mp model.MilestoneProgress
	err := r.db.Where("user_progress_id = ? AND milestone_id = ?", userProgressID, milestoneID).
		First(&mp).Error
	if err != nil {
		return nil, err
	}
	return &mp, nil
}

func (r *ProgressRepository) CreateMilestoneProgress(mp *model.MilestoneProgress) error {
	return r.db.Create(mp).Error
}

func (r *ProgressRepository) UpdateMilestoneProgress(mp *model.MilestoneProgress) error {
	return r.db.Save(mp).Error
}

func (r *ProgressRepository) FindResourceProgress(userProgressID, resourceID uint) (*model.ResourceProgress, error) {
	var rp model.ResourceProgress
	err := r.db.Where("user_progress_id = ? AND resource_id = ?", userProgressID, resourceID).
		First(&rp).Error
	if err != nil {
		return nil, err
	}
	return &rp, nil
}

func (r *ProgressRepository) CreateResourceProgress(rp *model.ResourceProgress) error {
	return r.db.Create(rp).Error
}

func (r *ProgressRepository) UpdateResourceProgress(rp *model.ResourceProgress) error {
	return r.db.Save(rp).Error
}

func (r *ProgressRepository) CountCompletedMilestones(userProgressID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.MilestoneProgress{}).
		Where("user_progress_id = ? AND status = ?", userProgressID, model.StatusCompleted).
		Count(&count).Error
	return count, err
}

// AssessmentRepository handles database operations for assessments.
type AssessmentRepository struct {
	db *gorm.DB
}

func NewAssessmentRepository(db *gorm.DB) *AssessmentRepository {
	return &AssessmentRepository{db: db}
}

func (r *AssessmentRepository) FindByMilestoneID(milestoneID uint) ([]model.SkillAssessment, error) {
	var assessments []model.SkillAssessment
	err := r.db.Where("milestone_id = ?", milestoneID).Find(&assessments).Error
	return assessments, err
}

func (r *AssessmentRepository) FindByID(id uint) (*model.SkillAssessment, error) {
	var a model.SkillAssessment
	err := r.db.First(&a, id).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AssessmentRepository) CreateResult(result *model.AssessmentResult) error {
	return r.db.Create(result).Error
}

func (r *AssessmentRepository) FindUserResults(userID uint) ([]model.AssessmentResult, error) {
	var results []model.AssessmentResult
	err := r.db.Where("user_id = ?", userID).Find(&results).Error
	return results, err
}

func (r *AssessmentRepository) FindMilestoneByID(id uint) (*model.Milestone, error) {
	var m model.Milestone
	err := r.db.First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}
