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
	//ini inisiasi userID yang mana ingin mendapatkan id si user
	getUserId := currentUser.ID

	newProduct, err := h.topUpService.CreateTopUp(input, getUserId)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, (newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *topUpHandler) GetTopUp(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	products, err := h.topUpService.GetTopUp((id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get product")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (products))
	c.JSON(http.StatusOK, response)
}
