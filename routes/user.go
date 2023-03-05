package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", middleware.Auth(h.FindUsers))
	e.GET("/user/:id", middleware.Auth(h.GetUserById))
	e.PATCH("/user/:id", middleware.Auth(h.UpdateUser))

}
