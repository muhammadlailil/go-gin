package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todos-apps/src/jwt_token"
)

func ApiMiddleware(c *gin.Context) {
	token := jwt_token.ExtractToken(c);
	expired,err := jwt_token.ExtractExpired(token)
	now := time.Now().Unix()
	if (token=="") {
		c.JSON(http.StatusUnauthorized,gin.H{
			"status" : "no_token",
		})
		c.Abort()
	}else if(err!=nil) {
		c.JSON(http.StatusUnauthorized,gin.H{
			"status" : "expired",
		})
		c.Abort()
	}else if(int64(expired)<now) {
		c.JSON(http.StatusUnauthorized,gin.H{
			"status" : "expired",
		})
		c.Abort()
	}
}
