package router

import (
	"ecom_clean_architecture/adapter/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/products", func(ctx echo.Context) error {
		return c.Product.GetAllProducts(ctx)
	})

	return e
}
