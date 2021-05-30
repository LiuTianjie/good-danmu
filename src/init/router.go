/*
 * @Descripttion: routers
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-23 19:55:32
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-29 10:51:18
 */
package init

import (
	"good-danmu/src/middleware"
	"good-danmu/src/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(Cors())
	PublicGroup := Router.Group("base")
	{
		router.InitBaseUser(PublicGroup)
	}
	PrivateGroup := Router.Group("v1")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitDanmuRoute(PrivateGroup)
		router.IintAuthRoute(PrivateGroup)
	}
	return Router
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
