/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 17:25:38
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-29 10:53:19
 */
package router

import (
	"good-danmu/src/api"
	"good-danmu/src/global"
	"good-danmu/src/handler"
	"good-danmu/src/model"
	"log"

	"github.com/gin-gonic/gin"
)

func InitDanmuRoute(Router *gin.RouterGroup) {
	DanmuRouter := Router.Group("")
	{
		DanmuRouter.GET("danmu:id", handler.DanmuHandler)
		DanmuRouter.GET("createdanmu", api.CreateDanmuRoom)
		DanmuRouter.GET("getdanmulist", api.GetDanmuRooms)
	}
}

func InitBaseUser(Router *gin.RouterGroup) {
	BaseUserRouter := Router.Group("")
	{
		BaseUserRouter.POST("register", api.Register)
		BaseUserRouter.POST("login", api.Login)
	}
}

func IintAuthRoute(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("")
	{
		AuthRouter.GET("user", api.UserInfo)
	}
}

func SaveDanmu(DanmuId, Username string, DanmuContent []byte) {
	danmu := &model.DanmuContent{
		Content:  DanmuContent,
		RoomId:   DanmuId,
		Username: Username,
		Type:     "normal",
	}
	err := global.DB.Create(&danmu).Error
	if err != nil {
		log.Println("插入出错")
	}
	log.Println("插入成功！")
}
