/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 17:25:38
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-23 19:47:38
 */
package router

import (
	"good-danmu/src/handler"

	"github.com/gin-gonic/gin"
)

func InitDanmuRoute(Router *gin.RouterGroup) {
	DanmuRouter := Router.Group("")
	{
		DanmuRouter.GET("danmu:id", handler.DanmuHandler)
	}
}
