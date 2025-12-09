package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/migrations"
	"gorm.io/gorm"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate.go [up|down]")
	}

	config.ConnectDB()
	db := config.DB

	if db == nil {
		panic("database is nil")
	}

	switch os.Args[1] {
	case "up":
		runMigrations(db)
	case "down":
		rollbackLast(db)
	case "status":
		checkStatus(db)
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}

func runMigrations(db *gorm.DB) {
	m := migrations.NewMigrator(db)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("✅ All migrations applied")
}

func rollbackLast(db *gorm.DB) {
	m := migrations.NewMigrator(db)
	if err := m.RollbackLast(); err != nil {
		log.Fatalf("Rollback failed: %v", err)
	}
	fmt.Println("✅ Last migration rolled back")
}

func checkStatus(db *gorm.DB) {
	var applied []string
	if err := db.Table("migrations").Pluck("id", &applied).Error; err != nil {
		log.Fatalf("Could not fetch migrations: %v", err)
	}

	all := migrations.GetMigrations()

	fmt.Println("== Migration Status ==")
	for _, m := range all {
		if slices.Contains(applied, m.ID) {
			fmt.Printf("[X] %s applied\n", m.ID)
		} else {
			fmt.Printf("[ ] %s pending\n", m.ID)
		}
	}
}
