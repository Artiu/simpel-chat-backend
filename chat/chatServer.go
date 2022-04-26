package controllers

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type messageReceived struct {
	ChatRoomId ChatRoomId
	Message    string
}

var rooms map[ChatRoomId]*ChatRoom = make(map[ChatRoomId]*ChatRoom)

func ConnectToRooms(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		chatRoomIds := []ChatRoomId{"1234", "Hello", "Czesc"}
		for _, id := range chatRoomIds {
			if _, ok := rooms[id]; !ok {
				rooms[id] = CreateChatRoom(id)
			}
			rooms[id].Connect(&Client{id: "gigamocnyzawodnik", connection: ws})
		}
		for {
			msg := messageReceived{}
			err := websocket.JSON.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				continue
			}
			rooms[msg.ChatRoomId].HandleMessage(Message{Message: msg.Message, UserId: "gigaMocnyZawodnik"})
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
