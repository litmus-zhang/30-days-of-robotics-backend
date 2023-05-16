package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuthenticated(c *gin.Context) {
	//cookie, _ := c.Cookie("30_DOR")
	session := sessions.Default(c)
	token := session.Get("token")

	if token == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}
	c.Next()
}

func GetUserId(c *gin.Context) interface{} {
	session := sessions.Default(c)
	id := session.Get("userID")

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return nil
	}
	return id
}
