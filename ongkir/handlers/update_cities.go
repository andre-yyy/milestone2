package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	entity "milestone2/entities"
	"milestone2/users"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func UpdateCities(c echo.Context, us *users.UserService) error {

	userId := c.Get("user_id").(string)
	user, err := us.GetUserById(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if user.Role != "admin" {
		response := entity.ServerResponse{
			Message: "this feature is only available for admin",
			Code:    http.StatusForbidden,
		}
		return c.JSON(http.StatusForbidden, response)
	}

	ctx := context.Background()

	url := "https://api.rajaongkir.com/starter/city"
	apiKey := os.Getenv("RAJAONGKIR_API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		response := entity.ServerResponse{
			Message: fmt.Sprintf("Error creating request : %v", err),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	req.Header.Set("key", apiKey)

	client := http.Client{}

	response, err := client.Do(req)
	if err != nil {
		response := entity.ServerResponse{
			Message: fmt.Sprintf("Error sending request : %v", err),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		response := entity.ServerResponse{
			Message: fmt.Sprintf("Error reading response : %v", err),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	var results entity.OngkirResponse[[]entity.City]

	err = json.Unmarshal(body, &results)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	cities, err := json.Marshal(results.RajaOngkir.Results)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	err = rdb.Set(ctx, "rajaongkircities", cities, 0).Err()
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	return c.JSON(http.StatusOK, results.RajaOngkir.Results)
}
