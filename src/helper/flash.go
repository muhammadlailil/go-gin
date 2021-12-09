package helper

import "github.com/gin-gonic/gin"

func SetFlash(c *gin.Context, name string, value string) {
	c.SetCookie(name, value, 10, "/", c.Request.URL.Hostname(), false, true)
}
func GetFlas(c *gin.Context, name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie
}
