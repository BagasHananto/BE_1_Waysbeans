package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.ControlCart(cartRepository)

	e.GET("/carts", h.FindCarts)
	e.GET("/cart/:id", h.GetCart)
	e.POST("/cart", middleware.Auth(h.CreateCart))
}
