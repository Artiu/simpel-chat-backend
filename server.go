package main

import (
	"simpel-chat/db"
	user "simpel-chat/user/register"
	validator "simpel-chat/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "gopkg.in/olahol/melody.v1"
)

func main() {
	e := echo.New()
	e.Validator = validator.New()
	db.Init()
	user.
		// m := melody.New()
		e.Use(middleware.Logger())
	e.POST("/register", user.RegisterHandler)
	// e.GET("/ws", func(c echo.Context) error {
	// 	m.HandleRequest(c.Response().Writer, c.Request())
	// 	return nil
	// })
	// m.HandleMessage(func(s *melody.Session, msg []byte) {
	// 	m.Broadcast(string(msg))
	// })
	e.Logger.Fatal(e.Start(":1323"))
}
