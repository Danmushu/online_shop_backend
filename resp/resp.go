package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"msg":  msg,
		"data": nil,
	})
	c.Abort()
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": data,
	})
	c.Abort()
}
