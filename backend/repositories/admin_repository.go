package repositories

import (
	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/models"
)

type AdminRepository interface {
	FindByEmail(email, password string) (*models.Admin, error)
}

type adminRepository struct{}

func AuthAdminRepository() AdminRepository {
	return &adminRepository{}
}

func (r *adminRepository) FindByEmail(email, password string) (*models.Admin, error) {
	var admin models.Admin
	if err := config.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
