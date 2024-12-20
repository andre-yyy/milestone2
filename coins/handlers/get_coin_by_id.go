package handlers

import (
	"net/http"

	"milestone2/coins"
	entity "milestone2/entities"
	"milestone2/topups"
	"milestone2/users"
	utility "milestone2/utilities"

	"github.com/labstack/echo/v4"
)

// @summary Get coin by id
// @description Get coin by id
// @tags Milestone 2
// @security Auth
// @produce json
// @param coinID path int true "coinID"
// @success 201 {object} entity.TopupInvoice[entity.Coin]
// @failure 400 {object} entity.ServerResponse
// @failure 401 {object} entity.ServerResponse
// @failure 404 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /coins/{coinID} [get]
func GetCoinById(c echo.Context, cs *coins.CoinService, ts topups.TopupService, us users.UserService) error {
	coinId := c.Param("id")
	userId := c.Get("user_id").(string)

	coin, err := cs.GetCoinById(coinId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	user, err := us.GetUserById(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	invoiceRequest := entity.InvoiceRequest{
		Name:        coin.Name,
		Price:       coin.Price * 1000,
		Description: coin.Description,
	}

	invoice, externalId, err := utility.CreateInvoice(invoiceRequest, user)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	topupToCreate := entity.Topup{
		ExternalID: externalId,
		UserID:     user.ID,
	}

	_, err = ts.CreateTopup(topupToCreate)

	response := entity.TopupInvoice[entity.Coin]{
		Message: "success create invoice",
		Code:    http.StatusCreated,
		Data:    coin,
		Invoice: *invoice,
	}
	return c.JSON(http.StatusCreated, response)
}
