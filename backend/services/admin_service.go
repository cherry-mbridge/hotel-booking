package services

import (
	"errors"

	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/repositories"
	adminService "github.com/cherry-mbridge/hotel-booking/backend/services/admin"
	"github.com/cherry-mbridge/hotel-booking/backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	repo repositories.AdminRepository
}

func AuthAdminService(repo repositories.AdminRepository) *AdminService {
	return &AdminService{repo}
}

func (s *AdminService) Login(c *gin.Context, email, password string) error {
	admin, err := s.repo.FindByEmail(email, password)

	if err != nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return errors.New("invalid email or password")
	}

	// create tokens
	accessToken, err := utils.CreateAccessToken(admin.ID, "admin", config.AccessTTL)
	if err != nil {
		return errors.New("couldn't create token")
	}
	refreshToken, err := utils.CreateAccessToken(admin.ID, "admin", config.RefreshTTL)
	if err != nil {
		return errors.New("couldn't create refresh token")
	}

	// Set cookies. For dev: secure=false; in production set secure=true and proper Domain.

	adminService.SetTokenCookie(c, "admin_token", accessToken, int(config.AccessTTL.Seconds()), false)

	adminService.SetTokenCookie(c, "refresh_token", refreshToken, int(config.RefreshTTL.Seconds()), false)
	return nil
}

func (s *AdminService) Logout(c *gin.Context) {
	// Remove cookies
	// c.SetCookie(config.AccessCookieName, "", -1,
	// 	"/", "localhost", false, true)

	// c.SetCookie(config.RefreshCookieName, "", -1,
	// 	"/", "localhost", false, true)

	adminService.SetTokenCookie(c, "admin_token", "", -1, false)
}

func (s *AdminService) Profile(c *gin.Context) (map[string]any, error) {
	token, err := c.Cookie("admin_token")
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
