package handler

import (
	"camera-rent/entity"
	"camera-rent/helper"
	"camera-rent/input"
	"camera-rent/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type topUpHandler struct {
	topUpService service.ServiceTopUp
}

func NewTopUpHandler(service service.ServiceTopUp) *topUpHandler {
	return &topUpHandler{service}
}

// CreatetopUp godoc
// @Summary Create a new top-up transaction
// @Description Create a top-up transaction for the current user with the provided amount
// @Accept json
// @Produce json
// @Tags TopUp
// @Param Authorization header string true "Bearer token"
// @Param body body input.InputTopUp true "Top-up details"
// @Success 200 {object} map[string]interface{} "Top-up successfully created"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 422 {object} map[string]interface{} "Validation error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/topup [post]
func (h *topUpHandler) CreatetopUp(c *gin.Context) {
	var input input.InputTopUp

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(*entity.User)
	// Inisiasi userID dari current user
	getUserId := currentUser.ID

	newTopUp, err := h.topUpService.CreateTopUp(input, getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, newTopUp)
	c.JSON(http.StatusOK, response)
}

// GetTopUp godoc
// @Summary Get top-up transaction details
// @Description Retrieve the details of a top-up transaction by its ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags TopUp
// @Param id path int true "Top-up Transaction ID"
// @Success 200 {object} map[string]interface{} "Top-up transaction details"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/topup/{id} [get]
func (h *topUpHandler) GetTopUp(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	topUp, err := h.topUpService.GetTopUp(id)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Error retrieving top-up")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, topUp)
	c.JSON(http.StatusOK, response)
}
