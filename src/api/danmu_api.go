/*
@Time : 2021/5/31 16:46
@Author : nickname4th
@File : danmu_api
@Software: GoLand
*/
package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"good-danmu/src/global"
	h "good-danmu/src/handler"
	"good-danmu/src/model"
	"good-danmu/src/utils"
	"good-danmu/src/utils/parse"
	"log"
)

type RoomStruct struct {
	RoomName string
	RoomDesc string
}

func GetDanmuRooms(c *gin.Context) {
	rooms := h.SearchAllRooms()
	//utils.OkMsg(200, "请求成功", c)
	ret, _ := json.Marshal(rooms)
	n := len(ret)        //Find the length of the byte array
	s := string(ret[:n]) //convert to string
	utils.OkDetail(200, s, "请求成功", c)
}

func CreateDanmuRoom(c *gin.Context) {
	//	TODO: use casbin to create more complex auth
	//	 here use my phone to simply control danmu room's creation.
	var (
		err      error
		Username string
		room     RoomStruct
		token    string
	)
	if err = c.ShouldBind(&room); err != nil {
		utils.FailedMsg(400, "参数错误", c)
	}
	token = c.GetHeader("token")
	Username, err = parse.GetUserFromToken(token)
	if err != nil || Username != "17330929598" {
		utils.FailedMsg(401, "权限不足", c)
		c.Abort()
		return
	} else {
		var roomInfo = model.DanmuRoom{
			RoomName: room.RoomName,
			RoomDesc: room.RoomDesc,
		}
		if err = global.DB.Create(&roomInfo).Error; err != nil {
			utils.FailedMsg(400, "创建错误", c)
			c.Abort()
			return
		} else {
			rooms := h.SearchAllRooms()
			utils.OkDetail(200, rooms, "创建成功！", c)
		}
	}
	log.Println("room", room)
}
