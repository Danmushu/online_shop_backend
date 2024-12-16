package main

import (
	"github.com/gin-gonic/gin"
	"project/global"
	"project/routers"
)

func main() {

	server := gin.Default()

	/** 路由 **/
	auth := server.Group("/auth")
	user := server.Group("/user")
	//todo
	order := server.Group("/order")
	item := server.Group("/item")

	routers.AuthenticationRoutes(auth)
	routers.UserRoutes(user)
	err := server.Run(global.API)
	if err != nil {
		return
	}
}
