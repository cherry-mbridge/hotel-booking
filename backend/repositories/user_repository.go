package repositories

import (
	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/models"
	"gorm.io/gorm"
)

// 1. Define the Interface (The Contract)
type UserRepository interface {
	FindByEmail(email string) (models.User, error)
	Create(user *models.User) error
	FindByID(id uint) (models.User, error)
}

// 2. Concrete Struct (The Implementation)
type userRepository struct { // Use lowercase for the concrete struct
	DB *gorm.DB
}

// 3. Constructor returns the Interface, hiding the concrete struct
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

// FindByEmail retrieves a user by their email address.
// 4. Attach methods to the concrete struct (which now satisfies the interface)
func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err // GORM will return gorm.ErrRecordNotFound
	}
	return user, nil
}

// Create creates a new user record in the database.
func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

// FindByID retrieves a user by their ID.
func (r *userRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}
