package handlers

import (
	"net/http"

	entity "milestone2/entities"
	"milestone2/topups"
	"milestone2/users"

	"github.com/labstack/echo/v4"
)

// @summary Get all topups
// @description Get all topups
// @tags Milestone 2
// @security Auth
// @produce json
// @Param status query string false "Status"
// @success 200 {object} entity.ServerResponseData[[]entity.Topup]
// @failure 401 {object} entity.ServerResponse
// @failure 404 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /topups [get]
func GetAllTopupsByStatus(c echo.Context, ts topups.TopupService, us users.UserService) error {
	status := c.QueryParam("status")
	userId := c.Get("user_id").(string)
	user, err := us.GetUserById(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if user.Role == "admin" {
		topups, err := ts.GetAllTopupsByStatus(status, nil)
		if err != nil {
			response := entity.ServerResponse{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
		response := entity.ServerResponseData[[]entity.Topup]{
			Message: "success get all topups.",
			Code:    http.StatusOK,
			Data:    topups,
		}
		return c.JSON(http.StatusOK, response)
	}
	topups, err := ts.GetAllTopupsByStatus(status, &userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := entity.ServerResponseData[[]entity.Topup]{
		Message: "success get all topups",
		Code:    http.StatusOK,
		Data:    topups,
	}
	return c.JSON(http.StatusOK, response)
}
