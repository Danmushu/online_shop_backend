package routers

import (
	"github.com/gin-gonic/gin"
	"project/controllers"
)

func SetUserRouter(api *gin.RouterGroup) {
	// 用户登录信息
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser) //注册
		auth.POST("", controllers.LoginUser)             //登录
		auth.POST("/logout", controllers.LogoutUser)     //退出
		auth.GET("", controllers.CurrentUser)
	}

	//商品
	product := api.Group("/product")
	{
		product.GET("", controllers.ProductGet)
		product.GET("/:id", controllers.ProductGet)
	}

	// 订单
	order := api.Group("/order")
	{
		order.POST("")
		order.GET("")
	}
}
