package services

import (
	"errors"

	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/models"
	"github.com/cherry-mbridge/hotel-booking/backend/repositories"
	userService "github.com/cherry-mbridge/hotel-booking/backend/services/user"
	"github.com/cherry-mbridge/hotel-booking/backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Register(c *gin.Context, name, email, password string) (models.User, error) {
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return models.User{}, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New("failed to hash password")
	}

	user := models.User{Name: name, Email: email, Password: string(hashedPassword)}
	if err := s.repo.Create(&user); err != nil {
		return models.User{}, errors.New("failed to create user")
	}
	return user, nil
}

func (s *UserService) Login(c *gin.Context, email, password string) error {
	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("invalid email or password")
	}

	// create tokens
	accessToken, err := utils.CreateAccessToken(user.ID, "user", config.AccessTTL)
	if err != nil {
		return errors.New("couldn't create token")
	}
	refreshToken, err := utils.CreateAccessToken(user.ID, "user", config.RefreshTTL)
	if err != nil {
		return errors.New("couldn't create refresh token")
	}

	// Set cookies. For dev: secure=false; in production set secure=true and proper Domain.

	userService.SetTokenCookie(c, "user_token", accessToken, int(config.AccessTTL.Seconds()), false)

	userService.SetTokenCookie(c, "refresh_token", refreshToken, int(config.RefreshTTL.Seconds()), false)
	return nil
}

func (s *UserService) Logout(c *gin.Context) {
	// Remove cookies
	userService.SetTokenCookie(c, "user_token", "", -1, false)
}

func (s *UserService) Profile(c *gin.Context) (map[string]any, error) {
	token, err := c.Cookie("user_token")
	if err != nil {
		return nil, errors.New("token error")
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		return nil, errors.New("token error")
	}
	return map[string]any{
		"type": claims["type"],
	}, nil
}
