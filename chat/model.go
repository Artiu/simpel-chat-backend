package chat

type Room struct {
	Id          string
	LastMessage Message
	Members     []Member
}

type Message struct {
	Id      string
	RoomID  string
	Content string
	Sender  string
}

type Member struct {
	Id       string
	Nick     string
	IsActive bool
}
