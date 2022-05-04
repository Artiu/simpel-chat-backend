package user

import (
	"context"
	"simpel-chat/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Nick          string   `bson:"_id"`
	Password      string   `bson:"password"`
	IsActive      bool     `bson:"is_active"`
	RefreshTokens []string `bson:"refresh_tokens"`
	JoinedRooms   []string `bson:"joined_rooms"`
}

func GetPasswordByNick(nick string) string {
	var user User
	db.Client.Database("simpel_chat").Collection("users").FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: nick}}, options.FindOne().SetProjection(bson.D{primitive.E{Key: "password", Value: 1}})).Decode(&user)
	return user.Password
}
