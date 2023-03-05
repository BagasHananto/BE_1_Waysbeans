package middleware

import (
	"net/http"
	"strings"
	dto "waysbeans/dto/result"
	jwtToken "waysbeans/pkg/jwt"

	"github.com/labstack/echo/v4"
)

// Declare Result struct here ...
type Result struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Create Auth function here ...
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: "Failed", Message: "unauthorized"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, Result{Status: "Failed", Message: "unathorized"})
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}
