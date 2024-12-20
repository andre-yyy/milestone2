package handlers

import (
	"fmt"
	"net/http"
	"strings"

	entity "milestone2/entities"
	"milestone2/users"
	dto "milestone2/users/dtos"
	utility "milestone2/utilities"

	"github.com/labstack/echo/v4"
)

// @summary Register user
// @description Register user
// @tags Milestone 2
// @security Auth
// @accept json
// @produce json
// @param user body dto.RegisterUserDto true "User"
// @success 201 {object} entity.ServerResponseData[entity.User]
// @failure 400 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /users/register [post]
func RegisterUser(c echo.Context, us *users.UserService) error {
	var userToRegister dto.RegisterUserDto
	err := c.Bind(&userToRegister)
	if err != nil {
		response := entity.ServerResponse{
			Message: "invalid json",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if userToRegister.FirstName == "" {
		response := entity.ServerResponse{
			Message: "first name is required",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if userToRegister.LastName == "" {
		response := entity.ServerResponse{
			Message: "last name is required",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if userToRegister.Email == "" {
		response := entity.ServerResponse{
			Message: "email is required",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if userToRegister.Password == "" {
		response := entity.ServerResponse{
			Message: "password is required",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	isEmailValid := utility.ValidateEmail(userToRegister.Email)
	if !isEmailValid {
		response := entity.ServerResponse{
			Message: "invalid email format",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	newUser := entity.User{
		FullName: fmt.Sprintf("%s %s", userToRegister.FirstName, userToRegister.LastName),
		Email:    userToRegister.Email,
		Password: userToRegister.Password,
	}

	registeredUser, err := us.RegisterUser(newUser)
	if err != nil {
		if strings.Contains(err.Error(), "failed") || strings.Contains(err.Error(), "registered") {
			response := entity.ServerResponse{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			}
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := entity.ServerResponse{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := entity.ServerResponseData[entity.User]{
		Message: "register success",
		Code:    http.StatusCreated,
		Data:    registeredUser,
	}
	return c.JSON(http.StatusCreated, response)

}
