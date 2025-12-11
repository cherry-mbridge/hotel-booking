package main

import (
	"time"

	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true, // ✅ cookie allow
		MaxAge:           12 * time.Hour,
	}))
	routes.SetupRouter(r)
	routes.UserRoutes(r)
	r.Run(":8080")
}
