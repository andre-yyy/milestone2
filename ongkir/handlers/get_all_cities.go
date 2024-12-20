package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	entity "milestone2/entities"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func GetAllCities(c echo.Context) error {
	ctx := context.Background()

	var results []entity.City

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	val, err := rdb.Get(ctx, "rajaongkircities").Result()
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = json.Unmarshal([]byte(val), &results)
	if err != nil {
		response := entity.ServerResponse{
			Message: "unmarshal error :" + err.Error(),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := entity.ServerResponseData[[]entity.City]{
		Message: "success get all cities",
		Code:    http.StatusOK,
		Data:    results,
	}
	return c.JSON(http.StatusOK, response)
}
