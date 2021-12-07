package restcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todos-apps/src/helper"
	"todos-apps/src/jwt_token"
)

//https://gowebexamples.com/password-hashing/
func PostLogin(c *gin.Context)  {

	email := c.PostForm("email")
	password := c.PostForm("password")
	//hash, _ := helper.HashMakePassword("12345678") // ignore error for the sake of simplicity


	hash := "$2a$14$tak1jX4YNmoa7mQYtR8GWOw2skm12i.0sXfYSbW2HNOEBLQNfyd.S"
	valid := helper.HashCheckPassword(password,hash);
	if (email=="admin@gmail.com" && valid) {

		data := &jwt_token.JwtTokenData{
			Scope :  "all",
			UsersId: 10,
			Email: "admin@gmail.com",
		}
		token,_ := jwt_token.GenerateToken(data)


		c.JSON(http.StatusOK,gin.H{
			"login" : "success",
			"token" : token,
		})
	}else{
		c.JSON(http.StatusUnauthorized,gin.H{
			"login" : "failed",
		})
	}
}
