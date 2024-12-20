package routes

import (
	"milestone2/products"
	product_handler "milestone2/products/handlers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductRoutes(mux *echo.Group, db *gorm.DB) *echo.Group {

	productRepository := products.NewProductRepository(db)
	productService := products.NewProductService(productRepository)

	mux.GET("", func(c echo.Context) error {
		return product_handler.GetAllProducts(c, productService)
	})

	mux.GET("/:id", func(c echo.Context) error {
		return product_handler.GetProductById(c, productService)
	})

	return mux
}
