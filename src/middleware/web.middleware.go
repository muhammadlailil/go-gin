package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebMiddleware(c *gin.Context) {
	//fmt.Println("hello from middleware")
	session := sessions.Default(c)
	users_id := session.Get("users_id")
	if users_id == nil {
		c.Redirect(http.StatusFound, "/login")
	}
}
