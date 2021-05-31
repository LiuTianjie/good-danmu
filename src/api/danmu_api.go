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
	"good-danmu/src/model"
	"good-danmu/src/utils"
	"good-danmu/src/utils/parse"
	"log"
)

func GetDanmuRooms(c *gin.Context) {
	type ans struct {
		DanmuId     string `json:"danmu_id"`
		RoomName    string `json:"room_name"`
		RoomDesc    string `json:"room_desc"`
		OnlineUsers int    `json:"online_users"`
		DanmuCount  int    `json:"danmu_count"`
	}
	var rooms []ans
	global.DB.Table("danmu_rooms").Select([]string{"danmu_id", "room_name", "room_desc", "online_users, danmu_count"}).Scan(&rooms)
	//utils.OkMsg(200, "请求成功", c)
	ret, _ := json.Marshal(rooms)
	n := len(ret)        //Find the length of the byte array
	s := string(ret[:n]) //convert to string
	utils.OkDetail(200, s, "请求成功", c)
}

func CreateDanmuRoom(c *gin.Context) {
	//	TODO: use casbin to create more complex auth
	//	 here use my phone to simply control danmu room's creation.
	token := c.GetHeader("token")
	room := c.ShouldBind(&model.DanmuRoom{})
	Username, err := parse.GetUserFromToken(token)
	if err != nil || Username != "17330929598" {
		utils.FailedMsg(401, "权限不足", c)
		c.Abort()
		return
	}
	log.Println(room)
}
