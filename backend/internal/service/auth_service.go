package service

import (
	"errors"

	"github.com/ronalc90/skillpath/internal/dto"
	"github.com/ronalc90/skillpath/internal/model"
	"github.com/ronalc90/skillpath/internal/repository"
	"github.com/ronalc90/skillpath/pkg/hash"
	jwtpkg "github.com/ronalc90/skillpath/pkg/jwt"
	"gorm.io/gorm"
)

var (
	ErrEmailAlreadyExists = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
)

// AuthService handles authentication business logic.
type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
	jwtExpiry int
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string, jwtExpiry int) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
	}
}

// Register creates a new user account.
func (s *AuthService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	existing, _ := s.userRepo.FindByEmail(req.Email)
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		DisplayName:  req.DisplayName,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := jwtpkg.GenerateToken(user.ID, user.Email, s.jwtSecret, s.jwtExpiry)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User:  toUserResponse(user),
	}, nil
}

// Login authenticates a user and returns a JWT token.
func (s *AuthService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if !hash.CheckPassword(req.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	token, err := jwtpkg.GenerateToken(user.ID, user.Email, s.jwtSecret, s.jwtExpiry)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User:  toUserResponse(user),
	}, nil
}

// GetProfile returns the authenticated user's profile.
func (s *AuthService) GetProfile(userID uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	resp := toUserResponse(user)
	return &resp, nil
}

// UpdateProfile updates the authenticated user's profile.
func (s *AuthService) UpdateProfile(userID uint, req dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	if req.DisplayName != "" {
		user.DisplayName = req.DisplayName
	}
	if req.AvatarURL != "" {
		user.AvatarURL = req.AvatarURL
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.GithubURL != "" {
		user.GithubURL = req.GithubURL
	}
	if req.LinkedinURL != "" {
		user.LinkedinURL = req.LinkedinURL
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	resp := toUserResponse(user)
	return &resp, nil
}

func toUserResponse(user *model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		AvatarURL:   user.AvatarURL,
		Bio:         user.Bio,
		GithubURL:   user.GithubURL,
		LinkedinURL: user.LinkedinURL,
	}
}
