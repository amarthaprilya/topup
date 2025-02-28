package handler

import (
	"camera-rent/formatter"
	"camera-rent/helper"
	"camera-rent/input"
	"camera-rent/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ServiceProduct
}

func NewProductHandler(service service.ServiceProduct) *productHandler {
	return &productHandler{service}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the given details
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Products
// @Param body body input.ProductInput true "Product details"
// @Success 200 {object} map[string]interface{} "Product successfully created"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 422 {object} map[string]interface{} "Validation error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/products [post]
func (h *productHandler) CreateProduct(c *gin.Context) {
	var input input.ProductInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, formatter.FormatterProduct(newProduct))
	c.JSON(http.StatusOK, response)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update an existing product by ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Products
// @Param id path int true "Product ID"
// @Param body body input.ProductInput true "Updated product details"
// @Success 200 {object} map[string]interface{} "Product successfully updated"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 422 {object} map[string]interface{} "Validation error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/products/{id} [put]
func (h *productHandler) UpdateProduct(c *gin.Context) {
	var input input.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getID := c.Param("id")
	param, _ := strconv.Atoi(getID)
	newProduct, err := h.productService.UpdateProduct(param, input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, newProduct)
	c.JSON(http.StatusOK, response)
}

// GetAllProduct godoc
// @Summary Get all products
// @Description Retrieve all products
// @Produce json
// @Tags Products
// @Success 200 {object} map[string]interface{} "List of products"
// @Failure 400 {object} map[string]interface{} "Error retrieving products"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/products [get]
func (h *productHandler) GetAllProduct(c *gin.Context) {
	products, err := h.productService.GetProducts()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Error retrieving products")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, products)
	c.JSON(http.StatusOK, response)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Retrieve a specific product using its ID
// @Produce json
// @Tags Products
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "Product details"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/products/{id} [get]
func (h *productHandler) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	product, err := h.productService.GetProduct(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Error retrieving product")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, product)
	c.JSON(http.StatusOK, response)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Security BearerAuth
// @Produce json
// @Security BearerAuth
// @Tags Products
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "Product successfully deleted"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/products/{id} [delete]
func (h *productHandler) DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := h.productService.DeleteProduct(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Error deleting product")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, "Product has been successfully deleted")
	c.JSON(http.StatusOK, response)
}
