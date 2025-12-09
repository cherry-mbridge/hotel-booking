package controllers

import (
	"net/http"

	"github.com/cherry-mbridge/hotel-booking/backend/dto"
	"github.com/cherry-mbridge/hotel-booking/backend/services"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service *services.AdminService
}

func AuthAdminController(service *services.AdminService) *AdminController {
	return &AdminController{service}
}

func (ctl *AdminController) Login(c *gin.Context) {
	var req dto.Admin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	err := ctl.service.Login(c, req.Email, req.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"type":    "admin",
	})
}

func (ctl *AdminController) Logout(c *gin.Context) {
	ctl.service.Logout(c)
	c.JSON(200, gin.H{"message": "Logged out"})
}

func (ctl *AdminController) Profile(c *gin.Context) {
	claims, err := ctl.service.Profile(c)
	if err != nil {
		// Not logged in
		c.JSON(200, gin.H{})
		return
	}

	c.JSON(200, gin.H{
		"type": claims["type"],
	})
}
