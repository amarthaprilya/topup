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

// will be called by user through payment endpoint
func (h *paymentSaldoHandler) DoPaymentSaldo(c *gin.Context) {
	makePaymentID := c.Param("id")

	var req input.SubmitPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	currentUser := c.MustGet("currentUser").(*entity.User)
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	resp, err := h.paymentSaldoService.DoPaymentSaldo(req, makePaymentID, getUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, (resp))
}

func (h *paymentSaldoHandler) GetPaymentSaldoNotification(c *gin.Context) {
	// orderID := c.Param("order_id")

	// params, err := strconv.Atoi(orderID)

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

	c.JSON(http.StatusOK, "succes")

	//1. get order data from db
	//2. check request transaction_status
	//3. map transaction_status to db payment status
	//4. update db payment status
}
