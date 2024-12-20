package handlers

import (
	"net/http"

	entity "milestone2/entities"
	"milestone2/topups"
	"milestone2/users"

	"github.com/labstack/echo/v4"
)

func GetNotification(c echo.Context, ts topups.TopupService, us users.UserService) error {

	var notification entity.Notification
	err := c.Bind(&notification)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	topup, err := ts.GetTopupByExternalId(notification.ExternalID)
	if err != nil {
		response := entity.ServerResponseData[entity.Notification]{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Data:    notification,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	topupToUpdate := entity.Topup{
		ExternalID: topup.ExternalID,
		Status:     notification.Status,
	}
	_, err = ts.UpdateTopupStatusByExternalId(topupToUpdate)
	if err != nil {
		response := entity.ServerResponseData[entity.Notification]{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Data:    notification,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	_, err = us.UpdateUserDepositById(topup.UserID, notification.Amount/1000)
	if err != nil {
		response := entity.ServerResponseData[entity.Notification]{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Data:    notification,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := entity.ServerResponseData[entity.Notification]{
		Message: "notified",
		Code:    http.StatusOK,
		Data:    notification,
	}
	return c.JSON(http.StatusOK, response)

}
