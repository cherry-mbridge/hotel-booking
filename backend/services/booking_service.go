package services

import (
	"github.com/cherry-mbridge/hotel-booking/backend/models"
	"github.com/cherry-mbridge/hotel-booking/backend/repositories"
)

type BookingService struct {
	repo repositories.BookingRepository
}

func NewBookingService(repo repositories.BookingRepository) *BookingService {
	return &BookingService{repo}
}

func (bs *BookingService) GetAll() ([]models.Booking, error) {
	return bs.repo.GetAll()
}

func (bs *BookingService) FindByID(id string) (*models.Booking, error) {
	return bs.repo.FindByID(id)
}
