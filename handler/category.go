package handler

import (
	"camera-rent/helper"
	"camera-rent/input"
	"camera-rent/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService service.ServiceCategory
}

func NewCategoryHandler(service service.ServiceCategory) *categoryHandler {
	return &categoryHandler{service}
}

// CreateCategory godoc
// @Summary Create new category
// @Description Create a new category with the provided name
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Categories
// @Param body body input.CategoryInput true "Category details"
// @Success 200 {object} map[string]interface{} "Category successfully created"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 422 {object} map[string]interface{} "Validation error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/category [post]
func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input input.CategoryInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCategory, err := h.categoryService.CreateCategory(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, newCategory)
	c.JSON(http.StatusOK, response)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update a category by ID
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Categories
// @Param id path int true "Category ID"
// @Param body body input.CategoryInput true "Updated category details"
// @Success 200 {object} map[string]interface{} "Category successfully updated"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 422 {object} map[string]interface{} "Validation error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/category/{id} [put]
func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var input input.CategoryInput

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
	updatedCategory, err := h.categoryService.UpdateCategory(param, input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, updatedCategory)
	c.JSON(http.StatusOK, response)
}

// GetAllCategory godoc
// @Summary Get all categories
// @Description Retrieve all categories
// @Produce json
// @Tags Categories
// @Success 200 {object} map[string]interface{} "List of categories"
// @Failure 400 {object} map[string]interface{} "Error retrieving categories"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/category [get]
func (h *categoryHandler) GetAllCategory(c *gin.Context) {
	categories, err := h.categoryService.GetCategorys()
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Error retrieving categories")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, categories)
	c.JSON(http.StatusOK, response)
}

// GetCategory godoc
// @Summary Get category by ID
// @Description Retrieve a specific category using its ID
// @Produce json
// @Tags Categories
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{} "Category details"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Category not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/category/{id} [get]
func (h *categoryHandler) GetCategory(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	category, err := h.categoryService.GetCategoryByID(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Error retrieving category")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, category)
	c.JSON(http.StatusOK, response)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete a category by ID
// @Security BearerAuth
// @Produce json
// @Security BearerAuth
// @Tags Categories
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{} "Category successfully deleted"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Category not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/category/{id} [delete]
func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	err := h.categoryService.DeleteCategory(int(id))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Error deleting category")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, "Category has been successfully deleted")
	c.JSON(http.StatusOK, response)
}
