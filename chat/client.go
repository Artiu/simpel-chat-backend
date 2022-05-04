package chat

import "golang.org/x/net/websocket"

type ClientId string

type Client struct {
	connection *websocket.Conn
	id         ClientId
}
