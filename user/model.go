package user

import "simpel-chat/db"

type User struct {
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required,min=10,max=100,hasLowercase,hasUppercase,hasNumber"`
	Nick          string
	RefreshTokens []string
	JoinedRooms   []string
}

func InitDBTable() {
	db.Session.Query(`CREATE TABLE IF NOT EXISTS (
		email text
		password text
	)`)
}
