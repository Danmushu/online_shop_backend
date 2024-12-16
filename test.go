package main

import (
	"github.com/gin-gonic/gin"
	"project/global"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is a test",
		})
	})

	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is a test",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "template/404.html", nil)
	})
	err := r.Run(global.API)
	if err != nil {
		return
	}

}
