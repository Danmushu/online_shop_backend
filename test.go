package main

import (
	"github.com/gin-gonic/gin"
	"project/global"
)

func main() {
	server := gin.Default()

	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is a test",
		})
	})

	server.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is a test",
		})
	})
	err := server.Run(global.API)
	if err != nil {
		return
	}

}
