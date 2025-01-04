package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"project/global"
	"project/middleware"
	"project/routers"
	"time"
)

func main() {

	// 日志记录
	_, err := os.Stat("./log/gin.log")
	if err == nil { // 文件存在
		now := time.Now()
		fn := "/" + now.Format("2006_01_02_150405") + ".log"
		err := os.Rename(global.LogDirPath+"/gin.log", global.LogDirPath+fn)
		if err != nil {
			fmt.Println("清除日志文件错误!")
			return
		}
	}
	log, err := os.Create("./log/gin.log")
	if err != nil {
		fmt.Println("日志输出设置错误")
		return
	}
	gin.DefaultWriter = io.MultiWriter(log, os.Stdout)

	server := gin.Default()
	server.Use(middleware.Cors)
	routers.SetWebsiteRouters(server)
	/** API路由 **/
	api := server.Group("/api")
	routers.SetUserRouter(api)

	err = server.Run(global.API)
	if err != nil {
		return
	}
}
