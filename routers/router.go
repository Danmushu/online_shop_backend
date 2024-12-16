package routers

import (
	"github.com/gin-gonic/gin"
	"project/controllers"
	"project/middleware"
)

func SetUserRouter(api *gin.RouterGroup) {
	// 用户登录信息
	auth := api.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
		auth.POST("/logout", controllers.LogoutUser)
		auth.GET("/info", controllers.CurrentUser)
	}

	api.Use(middleware.JwtAuth())

	//用户管理
	user := api.Group("/user")
	{
		user.GET("/detail", controllers.GetUserById)
		user.GET("/:id", controllers.GetUserById)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	//支付管理
	payment := api.Group("/payment")
	{
		payment.POST("/", controllers.CreatePayment)
		payment.GET("/", controllers.GetAllPayment)
		payment.GET("/:id", controllers.GetPaymentById)
		payment.PUT("/:id", controllers.UpdatePayment)
		payment.DELETE("/:id", controllers.DeletePayment)
	}

	// 购物车
	cart := api.Group("/cart")
	{
		cart.POST("/", controllers.CreateCart)
		cart.GET("/", controllers.GetAllCart)
		cart.GET("/:id", controllers.GetCartById)
		cart.PUT("/:id", controllers.UpdateCart)
		cart.DELETE("/:id", controllers.DeleteCart)
	}

	// 交易
	transaction := api.Group("/transaction")
	{
		transaction.POST("/", controllers.CreateTransaction)
		transaction.GET("/", controllers.GetAllTransaction)
		transaction.GET("/:id", controllers.GetTransactionById)
		transaction.PUT("/:id", controllers.UpdateTransaction)
		transaction.DELETE("/:id", controllers.DeleteTransaction)
	}

	//商品
	product := api.Group("/product")
	{
		product.POST("/", controllers.CreateProduct)
		product.GET("/", controllers.GetAllProduct)
		product.GET("/:id", controllers.GetProductById)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)

	}

	//商品种类
	category := api.Group("/category")
	{
		category.POST("/", controllers.CreateCategory)
		category.GET("/", controllers.GetAllCategory)
		category.GET("/:id", controllers.GetCategoryById)
		category.PUT("/:id", controllers.UpdateCategory)
		category.DELETE("/:id", controllers.DeleteCategory)

	}
}
