package handlers

import (
	"net/http"

	"milestone2/categories"
	entity "milestone2/entities"

	"github.com/labstack/echo/v4"
)

// @summary Get all categories
// @description Get all categories
// @tags Milestone 2
// @accept json
// @produce json
// @success 200 {object} entity.ServerResponseData[[]entity.Category]
// @failure 500 {object} entity.ServerResponse
// @router /categories [get]
func GetAllCategories(c echo.Context, cs *categories.CategoryService) error {
	allCategories, err := cs.GetAllCategories()
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}

		return c.JSON(http.StatusInternalServerError, response)
	}
	response := entity.ServerResponseData[[]entity.Category]{
		Message: "success get all categories",
		Code:    http.StatusOK,
		Data:    allCategories,
	}

	return c.JSON(http.StatusOK, response)
}
