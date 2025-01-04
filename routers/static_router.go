package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetWebsiteRouters(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/site/")
	})
	//server.Static("/cms/", "./websites/cms/") // 管理系统
}
