package restcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todos-apps/src/jwt_token"
)

func GenerateToken(ctx *gin.Context)  {

	data := &jwt_token.JwtTokenData{
		Scope :  "login",
	}
	token,_ := jwt_token.GenerateToken(data)
	dataToken,_ := jwt_token.Data(token)
	ctx.JSON(http.StatusOK,gin.H{
		"token" : token,
		"scode" : dataToken.Scope,
	})
}
