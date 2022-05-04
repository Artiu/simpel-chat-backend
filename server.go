package main

import (
	"log"
	"os"
	"simpel-chat/db"
	"simpel-chat/user"

	"simpel-chat/util/validator"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "gopkg.in/olahol/melody.v1"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file")
	}

	e := echo.New()
	e.Validator = validator.New()
	db.Init()
	defer db.Disconnect()
	// m := melody.New()
	e.Use(middleware.Logger())
	e.POST("/login", user.LoginHandler)
	e.POST("/register", user.RegisterHandler)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{TokenLookup: "cookie:accessToken", SigningKey: []byte(os.Getenv("JWT_SIGNING_KEY"))}))
	// e.GET("/ws", func(c echo.Context) error {
	// 	m.HandleRequest(c.Response().Writer, c.Request())
	// 	return nil
	// })
	// m.HandleMessage(func(s *melody.Session, msg []byte) {
	// 	m.Broadcast(string(msg))
	// })
	e.Logger.Fatal(e.Start(":1323"))
}
