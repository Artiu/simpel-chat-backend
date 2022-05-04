package chat

import (
	"encoding/json"
	"fmt"
)

type ChatRoomId string

//Maybe change to slice because of frequently send messages
type ChatRoom struct {
	clients map[ClientId]*Client
	id      ChatRoomId
}

func CreateChatRoom(id ChatRoomId) *ChatRoom {
	return &ChatRoom{
		clients: make(map[ClientId]*Client),
		id:      id,
	}
}

func (room *ChatRoom) Connect(client *Client) {
	room.clients[client.id] = client
	client.connection.Write([]byte(fmt.Sprintf("Connected to chat with id %s", room.id)))
}

func (room ChatRoom) HandleMessage(msg Message) {
	json, _ := json.Marshal(msg)
	for _, client := range room.clients {
		client.connection.Write(json)
	}
}
