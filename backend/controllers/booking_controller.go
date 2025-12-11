package controllers

import (
	"fmt"
	"net/http"

	"github.com/cherry-mbridge/hotel-booking/backend/services"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	service *services.BookingService
}

func NewBookingController(service *services.BookingService) *BookingController {
	return &BookingController{service}
}

func (bc *BookingController) GetAll(c *gin.Context) {
	bookings, err := bc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(bookings)
	// c.JSON(http.StatusOK, bookings)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bookings,
	})
}

func (bc *BookingController) Show(c *gin.Context) {
	id := c.Param("id")
	booking, err := bc.service.FindByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Booking not found"})
		return
	}
	c.JSON(200, booking)
}
