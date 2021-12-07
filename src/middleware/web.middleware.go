package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func WebMiddleware(*gin.Context) {
	fmt.Println("hello from middleware")
}
