package routes

import (
	"github.com/cherry-mbridge/hotel-booking/backend/controllers"
	middlewares "github.com/cherry-mbridge/hotel-booking/backend/middleware/admin"
	"github.com/cherry-mbridge/hotel-booking/backend/repositories"
	"github.com/cherry-mbridge/hotel-booking/backend/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Gin!"})
	})

	api_admin := r.Group("api/admin")
	{
		adminRepo := repositories.AuthAdminRepository()
		adminService := services.AuthAdminService(adminRepo)
		adminController := controllers.AuthAdminController(adminService)
		// Admin login
		api_admin.POST("login", adminController.Login)

		// Admin logout
		api_admin.GET("logout", adminController.Logout)

		// Admin me
		api_admin.Use(middlewares.AuthMiddleware())
		{
			api_admin.GET("me", adminController.Profile)

			bookingRepo := repositories.NewBookingRepository()
			bookingService := services.NewBookingService(bookingRepo)
			bookingController := controllers.NewBookingController(bookingService)

			api_admin.GET("bookings", bookingController.GetAll)
			api_admin.GET("bookings/:id", bookingController.Show)
		}
	}

}
