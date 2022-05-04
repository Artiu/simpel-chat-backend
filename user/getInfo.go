package user

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetInfoHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	return c.JSON(http.StatusOK, user.Claims)
}
