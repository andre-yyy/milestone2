package main

import (
	"log"
	"net/http"
	"os"

	"milestone2/coins"
	coin_handler "milestone2/coins/handlers"
	"milestone2/db"
	_ "milestone2/docs"
	milestone2middleware "milestone2/middlewares"
	notification_handler "milestone2/notifications/handlers"
	ongkir_handler "milestone2/ongkir/handlers"
	"milestone2/routes"
	"milestone2/topups"
	"milestone2/users"
	user_handler "milestone2/users/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Milestone 2 API
// @version 1.0
// @description Milestone 2 API
// @host ftgohacktiv8milestone2-3e32a9bf8f46.herokuapp.com
// @securityDefinitions.apiKey Auth
// @in header
// @name Authorization
// @description Authentication for Milestone 2 API
func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error load .env :", err)
		}
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)

	coinRepository := coins.NewCoinRepository(db)
	coinService := coins.NewCoinService(coinRepository)

	topupRepository := topups.NewTopupRepository(db)
	topupService := topups.NewTopupService(topupRepository)

	e.GET("/cities", func(c echo.Context) error {
		return ongkir_handler.GetAllCities(c)
	})

	e.GET("/cities/update", milestone2middleware.AuthMiddleware(func(c echo.Context) error {
		return ongkir_handler.UpdateCities(c, userService)
	}))

	e.GET("/coins", func(c echo.Context) error {
		return coin_handler.GetAllCoins(c, coinService)
	})

	e.GET("/coins/:id", milestone2middleware.AuthMiddleware(func(c echo.Context) error {
		return coin_handler.GetCoinById(c, coinService, *topupService, *userService)
	}))

	e.POST("/users/login", func(c echo.Context) error {
		return user_handler.LoginUser(c, userService)
	})

	e.POST("/users/register", func(c echo.Context) error {
		return user_handler.RegisterUser(c, userService)
	})

	e.POST("/notification", func(c echo.Context) error {
		return notification_handler.GetNotification(c, *topupService, *userService)
	})

	u := e.Group("/users")
	u.Use(milestone2middleware.AuthMiddleware)
	routes.UserRoutes(u, db, *userService)

	p := e.Group("/products")
	routes.ProductRoutes(p, db)

	o := e.Group("/orders")
	o.Use(milestone2middleware.AuthMiddleware)
	routes.OrderRoutes(o, db, *userService)

	t := e.Group("/topups")
	t.Use(milestone2middleware.AuthMiddleware)
	routes.TopupRoutes(t, db, *topupService, *userService)

	e.Logger.Fatal(e.Start(":" + port))
}
