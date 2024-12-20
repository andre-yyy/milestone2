package routes

import (
	"milestone2/topups"
	topup_handler "milestone2/topups/handlers"
	"milestone2/users"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TopupRoutes(mux *echo.Group, db *gorm.DB, topupService topups.TopupService, userService users.UserService) *echo.Group {

	mux.GET("", func(c echo.Context) error {
		return topup_handler.GetAllTopups(c, topupService, userService)
	})

	mux.GET("", func(c echo.Context) error {
		return topup_handler.GetAllTopupsByStatus(c, topupService, userService)
	})

	return mux
}
