package handler

import (
	"camera-rent/auth"
	"camera-rent/entity"
	"camera-rent/input"
	"camera-rent/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentSaldoHandler struct {
	paymentSaldoService service.ServicePaymentSaldo
	authService         auth.Service
}

func NewPaymentSaldoHandler(paymentSaldoService service.ServicePaymentSaldo, authService auth.Service) *paymentSaldoHandler {
	return &paymentSaldoHandler{paymentSaldoService, authService}
}

// DoPaymentSaldo godoc
// @Summary Process a top-up payment
// @Description Process a payment for a top-up transaction using the provided payment details and payment ID from the URL.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags PaymentSaldo
// @Param id path string true "Top-Up Payment ID"
// @Param Authorization header string true "Bearer token"
// @Param body body input.SubmitPaymentRequest true "Payment details"
// @Success 200 {object} entity.DoPayment "Successful payment response"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/payment/{id} [post]
func (h *paymentSaldoHandler) DoPaymentSaldo(c *gin.Context) {
	makePaymentID := c.Param("id")

	var req input.SubmitPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	currentUser := c.MustGet("currentUser").(*entity.User)
	// Inisiasi userID dari current user
	getUserId := currentUser.ID

	resp, err := h.paymentSaldoService.DoPaymentSaldo(req, makePaymentID, getUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetPaymentSaldoNotification godoc
// @Summary Process payment notification
// @Description Handle the notification from Midtrans regarding the top-up payment and update the payment status accordingly.
// @Accept json
// @Produce json
// @Tags PaymentSaldo
// @Param body body entity.MidtransNotificationRequest true "Midtrans notification payload"
// @Success 200 {string} string "top up was successful"
// @Failure 400 {object} map[string]interface{} "Invalid request payload"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/payment/notification [post]
func (h *paymentSaldoHandler) GetPaymentSaldoNotification(c *gin.Context) {
	var input *entity.MidtransNotificationRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	err := h.paymentSaldoService.HandleNotificationPaymentDonation(input)
	if err != nil {
		log.Printf("Error during payment: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "top up was successful")
}
