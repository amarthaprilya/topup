package handler

import (
	"camera-rent/entity"
	"camera-rent/input"
	"camera-rent/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	bookingService service.ServiceBooking
}

func NewBookingHandler(bookingService service.ServiceBooking) *BookingHandler {
	return &BookingHandler{bookingService}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var bookingInput input.BookingInput

	if err := c.ShouldBindJSON(&bookingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	newBooking, err := h.bookingService.CreateBooking(getUserId, bookingInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newBooking})
}

func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	bookings, err := h.bookingService.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

func (h *BookingHandler) GetBookingById(c *gin.Context) {
	getID := c.Param("id")
	param, _ := strconv.Atoi(getID)

	booking, err := h.bookingService.GetBookingById(param)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// func (h *BookingHandler) UpdateBooking(c *gin.Context) {
// 	getID := c.Param("id")
// 	param, _ := strconv.Atoi(getID)
// 	var bookingInput input.BookingInput

// 	if err := c.ShouldBindJSON(&bookingInput); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	updatedBooking, err := h.bookingService.UpdateBooking(id, entity.Booking{
// 		FirstDateRent: bookingInput.FirstDateRent,
// 		LastDateRent:  bookingInput.LastDateRent,
// 		ProductID:     bookingInput.ProductID,
// 		UserID:        bookingInput.UserID,
// 	})

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": updatedBooking})
// }

func (h *BookingHandler) DeleteBooking(c *gin.Context) {
	getID := c.Param("id")
	param, _ := strconv.Atoi(getID)

	err := h.bookingService.DeleteBooking(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
