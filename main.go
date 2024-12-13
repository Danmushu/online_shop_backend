package main

import (
	"github.com/gin-gonic/gin"
	"project/global"
)

func main() {

	server := gin.Default()

	err := server.Run(global.API)
	if err != nil {
		return
	}
}
