package seeders

import (
	"fmt"
	"log"

	"github.com/cherry-mbridge/hotel-booking/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AdminSeeder(db *gorm.DB) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	admin := models.Admin{
		Name:     "admin",
		Email:    "admin@email.com",
		Password: string(passwordHash),
	}

	if err := db.FirstOrCreate(&admin, models.Admin{Email: admin.Email}).Error; err != nil {
		log.Fatalf("Failed to seed admin: %v", err)
	}

	fmt.Println("✅ Admin seeded:", admin.Email)
}
