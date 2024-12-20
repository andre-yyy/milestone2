package routes

import (
	"milestone2/activities"
	activity_handler "milestone2/activities/handlers"
	"milestone2/users"
	user_handler "milestone2/users/handlers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(mux *echo.Group, db *gorm.DB, userService users.UserService) *echo.Group {

	activityRepository := activities.NewActivityRepository(db)
	activityService := activities.NewActivityService(activityRepository)

	mux.GET("", func(c echo.Context) error {
		return user_handler.GetUserById(c, &userService)
	})

	mux.GET("/activities", func(c echo.Context) error {
		return activity_handler.GetAllActivities(c, *activityService, userService)
	})

	mux.PATCH("", func(c echo.Context) error {
		return user_handler.UpdateUserLocationById(c, &userService)
	})

	return mux
}
