package main

import (
	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/seeders"
)

func main() {
	config.ConnectDB()
	db := config.DB
	seeders.AdminSeeder(db)
}
