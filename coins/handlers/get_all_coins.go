package handlers

import (
	"net/http"

	"milestone2/coins"
	entity "milestone2/entities"

	"github.com/labstack/echo/v4"
)

// @summary Get all coins
// @description Get all coins
// @tags Milestone 2
// @accept json
// @produce json
// @success 200 {object} entity.ServerResponseData[[]entity.Coin]
// @failure 500 {object} entity.ServerResponse
// @router /coins [get]
func GetAllCoins(c echo.Context, cs *coins.CoinService) error {
	allCoins, err := cs.GetAllCoins()
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := entity.ServerResponseData[[]entity.Coin]{
		Message: "success get all coins",
		Code:    http.StatusOK,
		Data:    allCoins,
	}
	return c.JSON(http.StatusOK, response)
}
