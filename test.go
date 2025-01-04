package main

import (
	"fmt"
)

func main() {
	//r := gin.Default()
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}
	fmt.Println(s)
	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "this is a test",
	//	})
	//})
	//
	//r.POST("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "this is a test",
	//	})
	//})
	//
	//r.NoRoute(func(c *gin.Context) {
	//	c.HTML(404, "template/404.html", nil)
	//})
	//err = r.Run(global.API)
	//if err != nil {
	//	return
	//}

}
