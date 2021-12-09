package logincontroller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"todos-apps/src/helper"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login.index.html", gin.H{
		"page_title": "Login",
	})
}

func PostLogin(c *gin.Context) {
	var (
		email    = c.PostForm("email")
		password = c.PostForm("password")
	)

	hash := "$2a$14$Bi3klIuhSK5swYZjc1MWze9xsENz7FUr5Dy3I1JU0q9I9BOZFNLCi"
	valid := helper.HashCheckPassword(password, hash)
	if email == "laililmahvut@gmail.com" && valid {
		session := sessions.Default(c)
		session.Set("users_id", 1)
		session.Set("name", "Muhammad Lailil Mahfud")
		_ = session.Save()

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Akun tidak ditemukan",
		})
	}
}

func GetLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
	c.Redirect(http.StatusFound, "/login")
}
