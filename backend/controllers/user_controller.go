package controllers

import (
	"net/http"

	"github.com/cherry-mbridge/hotel-booking/backend/dto"
	"github.com/cherry-mbridge/hotel-booking/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}

func (ctl *UserController) Register(c *gin.Context) {
	var req dto.RegisterUserDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		errs := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)

		for _, e := range errs {
			field := e.Field()

			switch e.Tag() {
			case "required":
				errorMessages[field] = field + " is required"
			case "email":
				errorMessages[field] = "Invalid email format"
			case "min":
				errorMessages[field] = field + " must be at least " + e.Param() + " characters"
			}
		}

		c.JSON(400, gin.H{"errors": errorMessages})
		return
	}

	// 2. Access the service via the receiver 'c'
	user, err := ctl.UserService.Register(c, req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
		"user":    gin.H{"id": user.ID, "name": user.Name},
		"type":    "user",
	})
}

func (ctl *UserController) Login(c *gin.Context) {
	var req dto.LoginUserDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	err := ctl.UserService.Login(c, req.Email, req.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"type":    "user",
	})
}

func (ctl *UserController) Logout(c *gin.Context) {
	ctl.UserService.Logout(c)
	c.JSON(200, gin.H{"message": "Logged out"})
}

func (ctl *UserController) Profile(c *gin.Context) {
	claims, err := ctl.UserService.Profile(c)
	if err != nil {
		// Not logged in
		c.JSON(200, gin.H{})
		return
	}

	c.JSON(200, gin.H{
		"type": claims["type"],
	})
}
