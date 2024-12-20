package routes

import (
	"milestone2/orders"
	order_handler "milestone2/orders/handlers"
	"milestone2/products"
	"milestone2/users"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func OrderRoutes(mux *echo.Group, db *gorm.DB, userService users.UserService) *echo.Group {

	orderRepository := orders.NewOrderRepository(db)
	orderService := orders.NewOrderService(orderRepository)

	productRepository := products.NewProductRepository(db)
	productService := products.NewProductService(productRepository)

	mux.GET("", func(c echo.Context) error {
		return order_handler.GetAllOrders(c, *orderService, userService)
	})

	mux.GET("/overdue", func(c echo.Context) error {
		return order_handler.GetAllOverdueOrders(c, *orderService, userService)
	})

	mux.POST("", func(c echo.Context) error {
		return order_handler.CreateOrder(c, *orderService, *productService, userService)
	})

	mux.PATCH("/:id", func(c echo.Context) error {
		return order_handler.UpdateOrderStatusById(c, orderService, userService)
	})

	return mux
}
