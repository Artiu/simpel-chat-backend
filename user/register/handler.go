package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10,max=100,hasLowercase,hasUppercase,hasNumber"`
}

func RegisterHandler(c echo.Context) error {
	registerData := new(RegisterRequest)
	if err := c.Bind(registerData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(registerData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusCreated)
}
