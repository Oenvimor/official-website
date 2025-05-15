package middlewares

import (
	"cqupt_hub/controller"
	"cqupt_hub/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URL
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxx.xxx.xxx
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割 判断格式是否有误
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 && parts[0] != "Bearer" {
			controller.ResponseError(c, controller.CodeInvalidAuth)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，使用之前定义好的解析JWT的函数来解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidAuth)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文中
		c.Set(controller.CtxUsernameKey, mc.Username)
		c.Next() // 后续的处理函数可以用c.Get("username")来获取当前请求的用户信息
	}
}
