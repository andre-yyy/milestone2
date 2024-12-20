package handlers

import (
	"fmt"
	"net/http"

	entity "milestone2/entities"
	"milestone2/users"
	dto "milestone2/users/dtos"

	"github.com/labstack/echo/v4"
)

func UpdateUserLocationById(c echo.Context, us *users.UserService) error {
	var userToUpdate dto.UpdateUserLocationDto
	userId := c.Get("user_id").(string)
	err := c.Bind(&userToUpdate)
	if err != nil {
		response := entity.ServerResponse{
			Message: "invalid json",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := us.GetUserById(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if userToUpdate.Address == "" {
		response := entity.ServerResponse{
			Message: fmt.Sprintf("address must not be empty"),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if userToUpdate.CityID == "" {
		response := entity.ServerResponse{
			Message: fmt.Sprintf("city id must not be empty"),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	user.Address = userToUpdate.Address
	user.CityID = userToUpdate.CityID

	updatedUser, err := us.UpdateUserLocationById(user)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := entity.ServerResponseData[entity.User]{
		Message: "success update user location",
		Code:    http.StatusOK,
		Data:    updatedUser,
	}
	return c.JSON(http.StatusOK, response)

}
