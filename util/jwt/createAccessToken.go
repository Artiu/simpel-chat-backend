package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserJWTClaims struct {
	Nick string `json:"nick"`
	jwt.StandardClaims
}

func CreateAccessToken(nick string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserJWTClaims{nick, jwt.StandardClaims{ExpiresAt: time.Now().Add(24 * time.Hour).Unix()}})
	signedString, _ := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	return signedString
}
