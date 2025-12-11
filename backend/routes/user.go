package routes

import (
	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/cherry-mbridge/hotel-booking/backend/controllers"
	middlewares "github.com/cherry-mbridge/hotel-booking/backend/middleware/user"
	"github.com/cherry-mbridge/hotel-booking/backend/repositories"
	"github.com/cherry-mbridge/hotel-booking/backend/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	db := config.DB
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	api_user := r.Group("api/users")
	{
		api_user.POST("/register", userController.Register)
		api_user.POST("/login", userController.Login)

		api_user.Use(middlewares.UserAuthMiddleware())
		{
			api_user.GET("/me", userController.Profile)
			api_user.POST("/logout", userController.Logout)
		}
	}
}
