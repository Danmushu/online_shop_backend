package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	//如果Origin不为空，说明这是一个跨域请求，因此设置了一系列CORS相关的响应头
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*")                                                                                                                          // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")                                                                                    //允许的HTTP方法
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")                                                             //允许的HTTP请求头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type") //可以被前端JavaScript读取的响应头
		c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                  //是否允许发送Cookie
	}
	//对于OPTIONS请求（预检请求），直接返回204 No Content状态码并终止请求处理
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}
