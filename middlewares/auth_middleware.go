package milestone2middleware

import (
	"net/http"
	"strings"

	entity "milestone2/entities"
	utility "milestone2/utilities"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			response := entity.ServerResponse{
				Message: "Please provide a token in the request headers for authentication",
				Code:    http.StatusUnauthorized,
			}
			return c.JSON(http.StatusUnauthorized, response)
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		if tokenString == "" {
			response := entity.ServerResponse{
				Message: "Please provide a token in the request headers for authentication",
				Code:    http.StatusUnauthorized,
			}
			return c.JSON(http.StatusUnauthorized, response)
		}

		token, err := utility.ParseToken(tokenString)
		if err != nil {
			response := entity.ServerResponse{
				Message: "Please provide a token in the request headers for authentication",
				Code:    http.StatusUnauthorized,
			}
			return c.JSON(http.StatusUnauthorized, response)
		}
		c.Set("user_id", token.Subject)
		return next(c)
	}
}
