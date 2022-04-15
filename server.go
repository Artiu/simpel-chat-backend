package main

import (
	controllers "simpel-chat/controllers/chat"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/chat", controllers.ConnectToRooms)
	e.Logger.Fatal(e.Start(":1323"))
}
