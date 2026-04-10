package service

import (
	"errors"
	"time"

	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/model"
	"github.com/ronalc90/skillpath/internal/repository"
)

var (
	ErrNoAssessment     = errors.New("no assessment found for this milestone")
	ErrInvalidAnswers   = errors.New("answers do not match the assessment questions")
	ErrMilestoneNotFound = errors.New("milestone not found")
)

const passingScorePercentage = 70

// AssessmentService handles skill assessment business logic.
type AssessmentService struct {
	assessmentRepo *repository.AssessmentRepository
}

func NewAssessmentService(assessmentRepo *repository.AssessmentRepository) *AssessmentService {
	return &AssessmentService{assessmentRepo: assessmentRepo}
}

// GetAssessment retrieves quiz questions for a milestone (without correct answers).
func (s *AssessmentService) GetAssessment(milestoneID uint) ([]dto.AssessmentQuestionResponse, error) {
	assessments, err := s.assessmentRepo.FindByMilestoneID(milestoneID)
	if err != nil {
		return nil, err
	}
	if len(assessments) == 0 {
		return nil, ErrNoAssessment
	}

	questions := make([]dto.AssessmentQuestionResponse, 0, len(assessments))
	for _, a := range assessments {
		questions = append(questions, dto.AssessmentQuestionResponse{
			ID:       a.ID,
			Question: a.Question,
			Options:  []string(a.Options),
		})
	}

	return questions, nil
}

// SubmitAssessment evaluates the user's answers and stores the result.
func (s *AssessmentService) SubmitAssessment(userID, milestoneID uint, req dto.SubmitAssessmentRequest) (*dto.AssessmentResultResponse, error) {
	assessments, err := s.assessmentRepo.FindByMilestoneID(milestoneID)
	if err != nil {
		return nil, err
	}
	if len(assessments) == 0 {
		return nil, ErrNoAssessment
	}

	questionMap := make(map[uint]*model.SkillAssessment)
	for i := range assessments {
		questionMap[assessments[i].ID] = &assessments[i]
	}

	score := 0
	details := make([]dto.AnswerDetail, 0, len(req.Answers))

	for _, answer := range req.Answers {
		question, ok := questionMap[answer.QuestionID]
		if !ok {
			continue
		}

		isCorrect := answer.SelectedIndex == question.CorrectAnswer
		if isCorrect {
			score++
		}

		details = append(details, dto.AnswerDetail{
			QuestionID:    answer.QuestionID,
			IsCorrect:     isCorrect,
			CorrectAnswer: question.CorrectAnswer,
			Explanation:   question.Explanation,
		})
	}

	totalQuestions := len(assessments)
	passed := float64(score)/float64(totalQuestions)*100 >= float64(passingScorePercentage)

	result := &model.AssessmentResult{
		UserID:         userID,
		MilestoneID:    milestoneID,
		Score:          score,
		TotalQuestions: totalQuestions,
	}

	if passed {
		now := time.Now()
		result.PassedAt = &now
	}

	if err := s.assessmentRepo.CreateResult(result); err != nil {
		return nil, err
	}

	return &dto.AssessmentResultResponse{
		Score:          score,
		TotalQuestions: totalQuestions,
		Passed:         passed,
		Details:        details,
	}, nil
}
