package user

import (
	"github.com/cherry-mbridge/hotel-booking/backend/config"
	"github.com/gin-gonic/gin"
)

// Helper to set cookie
func SetTokenCookie(c *gin.Context, name, token string, maxAge int, secure bool) {
	c.SetCookie(
		"access_token",
		token,
		int(config.AccessTTL.Seconds()),
		"/",
		"localhost",
		false, // secure=false for http
		true,  // httpOnly
	)

	// append SameSite
	c.Header("Set-Cookie", c.Writer.Header().Get("Set-Cookie")+"; SameSite=Lax")

}
