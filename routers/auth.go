package routers

import (
	"github.com/gin-gonic/gin"
	"project/controllers"
)

func AuthenticationRoutes(r *gin.RouterGroup) {
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/logout", controllers.LogoutUser)
	r.GET("/info", controllers.CurrentUser)
}
