package routers

import (
	"github.com/gin-gonic/gin"
	"project/controllers"
)

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/detail", controllers.GetUserById)
	r.GET("/:id", controllers.GetUserById)
	r.PUT("/:id", controllers.UpdateUser)
	r.DELETE("/:id", controllers.DeleteUser)
}
