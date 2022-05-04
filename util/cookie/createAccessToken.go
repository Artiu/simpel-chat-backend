package cookie

import (
	"net/http"
	"time"
)

func CreateAccessToken(accessToken string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "accessToken"
	cookie.Value = accessToken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// cookie.HttpOnly = true
	// cookie.SameSite = http.SameSiteStrictMode
	// cookie.Secure = true
	// cookie.Domain = "admin.artiu.com"
	return cookie
}
