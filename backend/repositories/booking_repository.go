package repositories

import (
	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/models"

	"gorm.io/gorm"
)

type BookingRepository interface {
	GetAll() ([]models.Booking, error)
	FindByID(id string) (*models.Booking, error)
}

type bookingRepository struct {
	DB *gorm.DB
}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{DB: config.DB}
}

func (br *bookingRepository) GetAll() ([]models.Booking, error) {
	var bookings []models.Booking
	if err := br.DB.Find(&bookings).Error; err != nil {
		return nil, err
	}

	return bookings, nil
}

func (br *bookingRepository) FindByID(id string) (*models.Booking, error) {
	var booking models.Booking

	if err := br.DB.First(&booking, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &booking, nil
}
