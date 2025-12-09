package migrations

import (
	"github.com/cherry-mbridge/hotel-booking/backend/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// CreateAdmins creates "admins" table
var CreateAdmins = &gormigrate.Migration{
	ID: "20251030_create_admins",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&models.Admin{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("admins")
	},
}
