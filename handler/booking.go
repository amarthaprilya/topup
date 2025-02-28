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

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a booking for the current user with the provided booking details.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Bookings
// @Param Authorization header string true "Bearer token"
// @Param booking body input.BookingInput true "Booking details"
// @Success 201 {object} map[string]interface{} "Booking successfully created"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/booking [post]
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var bookingInput input.BookingInput

	if err := c.ShouldBindJSON(&bookingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	currentUser := c.MustGet("currentUser").(*entity.User)
	// Inisiasi userID dari current user
	getUserId := currentUser.ID

	newBooking, err := h.bookingService.CreateBooking(getUserId, bookingInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newBooking})
}

// GetAllBookings godoc
// @Summary Get all bookings
// @Description Retrieve a list of all bookings.
// @Produce json
// @Tags Bookings
// @Success 200 {object} map[string]interface{} "List of bookings"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/booking [get]
func (h *BookingHandler) GetAllBookings(c *gin.Context) {
	bookings, err := h.bookingService.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// GetBookingById godoc
// @Summary Get booking by ID
// @Description Retrieve a booking by its ID.
// @Produce json
// @Tags Bookings
// @Param id path int true "Booking ID"
// @Success 200 {object} map[string]interface{} "Booking details"
// @Failure 404 {object} map[string]interface{} "Booking not found"
// @Router /api/booking/{id} [get]
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

// DeleteBooking godoc
// @Summary Delete booking
// @Description Delete a booking by its ID.
// @Produce json
// @Security BearerAuth
// @Tags Bookings
// @Param id path int true "Booking ID"
// @Success 200 {object} map[string]interface{} "Booking deleted successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/booking/{id} [delete]
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
