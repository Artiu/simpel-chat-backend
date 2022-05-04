package user

import (
	"context"
	"net/http"
	"simpel-chat/db"
	"simpel-chat/util/cookie"
	"simpel-chat/util/jwt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Nick     string `validate:"required,min=5,max=30"`
	Password string `validate:"required,min=10,max=100,hasLowercase,hasUppercase,hasNumber"`
}

func RegisterHandler(c echo.Context) error {
	registerData := new(RegisterRequest)
	if err := c.Bind(registerData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(registerData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	_, err := db.Client.Database("simpel_chat").Collection("users").InsertOne(context.TODO(), User{Nick: registerData.Nick, Password: string(hashedPassword)})
	if err != nil {
		return c.String(http.StatusConflict, err.Error())
	}
	accessToken := jwt.CreateAccessToken(registerData.Nick)
	cookie := cookie.CreateAccessToken(accessToken)
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
