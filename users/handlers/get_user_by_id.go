package handlers

import (
	"net/http"

	entity "milestone2/entities"
	"milestone2/users"

	"github.com/labstack/echo/v4"
)

// @summary Login user
// @description Login user
// @tags Milestone 2
// @security Auth
// @produce json
// @success 200 {object} entity.ServerResponseData[entity.User]
// @failure 401 {object} entity.ServerResponse
// @failure 404 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /users/register [post]
func GetUserById(c echo.Context, us *users.UserService) error {
	userId := c.Get("user_id").(string)
	user, err := us.GetUserById(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	response := entity.ServerResponseData[entity.User]{
		Message: "success get user data",
		Code:    http.StatusOK,
		Data:    user,
	}
	return c.JSON(http.StatusOK, response)
}
