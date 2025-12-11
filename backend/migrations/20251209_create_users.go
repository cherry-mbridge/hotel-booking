package migrations

import (
	"github.com/cherry-mbridge/hotel-booking/backend/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// CreateUsers creates "users" table
var CreateUsers = &gormigrate.Migration{
	ID: "20251209_create_users",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&models.User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
