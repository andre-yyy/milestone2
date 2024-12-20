package handlers

import (
	"net/http"
	"strings"

	entity "milestone2/entities"
	"milestone2/users"
	dto "milestone2/users/dtos"

	"github.com/labstack/echo/v4"
)

// @summary Login user
// @description Login user
// @tags Milestone 2
// @security Auth
// @accept json
// @produce json
// @param user body dto.LoginUserDto true "User"
// @success 200 {object} entity.ServerResponseData[entity.User]
// @failure 400 {object} entity.ServerResponse
// @failure 401 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /users/login [post]
func LoginUser(c echo.Context, us *users.UserService) error {
	var userToLogin dto.LoginUserDto
	err := c.Bind(&userToLogin)
	if err != nil {
		response := entity.ServerResponse{
			Message: "invalid json",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if userToLogin.Email == "" {
		response := entity.ServerResponse{
			Message: "email is required",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if userToLogin.Password == "" {
		response := entity.ServerResponse{
			Message: "password is required",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	loggedInUser, tokenString, err := us.LoginUser(userToLogin)
	if err != nil {
		if strings.Contains(err.Error(), "password") || strings.Contains(err.Error(), "not found") {
			response := entity.ServerResponse{
				Message: err.Error(),
				Code:    http.StatusUnauthorized,
			}
			return c.JSON(http.StatusUnauthorized, response)
		} else {
			response := entity.ServerResponse{
				Message: "internal server error",
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := entity.LoginSuccess[entity.User]{
		Message: "login success",
		Code:    http.StatusOK,
		Data:    loggedInUser,
		Token:   tokenString,
	}
	return c.JSON(http.StatusOK, response)

}
