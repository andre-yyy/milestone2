package handlers

import (
	"net/http"

	"milestone2/activities"
	entity "milestone2/entities"
	"milestone2/users"

	"github.com/labstack/echo/v4"
)

// @summary Get all user activities
// @description Get all user activities
// @tags Milestone 2
// @security Auth
// @produce json
// @success 200 {object} entity.ServerResponseData[[]entity.Activity]
// @failure 401 {object} entity.ServerResponse
// @failure 404 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /users/activities [get]
func GetAllActivities(c echo.Context, as activities.ActivityService, us users.UserService) error {
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
		activities, err := as.GetAllActivities()
		if err != nil {
			response := entity.ServerResponse{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
		response := entity.ServerResponseData[[]entity.Activity]{
			Message: "Success get all activities.",
			Code:    http.StatusOK,
			Data:    activities,
		}
		return c.JSON(http.StatusOK, response)
	}
	activities, err := as.GetAllActivitiesByUserId(user.ID)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := entity.ServerResponseData[[]entity.Activity]{
		Message: "success get all activities",
		Code:    http.StatusOK,
		Data:    activities,
	}
	return c.JSON(http.StatusOK, response)
}
