package migrations

import (
	"github.com/cherry-mbridge/hotel-booking/backend/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var CreateBookings = &gormigrate.Migration{
	ID: "20251030_create_bookings",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&models.Booking{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("bookings")
	},
}
