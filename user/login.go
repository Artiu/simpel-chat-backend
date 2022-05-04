package user

import (
	"log"
	"net/http"
	"simpel-chat/util/cookie"
	"simpel-chat/util/jwt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Nick     string `validate:"required,min=5,max=30"`
	Password string `validate:"required,min=10,max=100"`
}

func LoginHandler(c echo.Context) error {
	loginData := new(LoginRequest)
	if err := c.Bind(loginData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(loginData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	password := GetPasswordByNick(loginData.Nick)
	log.Println(password)
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginData.Password)); err != nil {
		return c.String(http.StatusUnauthorized, "Credentials are not correct")
	}
	accessToken := jwt.CreateAccessToken(loginData.Nick)
	cookie := cookie.CreateAccessToken(accessToken)
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
