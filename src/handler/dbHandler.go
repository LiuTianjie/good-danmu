/*
Package handler
@Time : 2021/6/3 16:50
@Author : nickname4th
@File : dbHandler
@Software: GoLand
*/
package handler

import (
	"good-danmu/src/global"
	"good-danmu/src/model"
	"good-danmu/src/utils"
)

// 返回房间信息结构体
type room struct {
	ID          string `json:"id"`
	RoomName    string `json:"room_name"`
	RoomDesc    string `json:"room_desc"`
	OnlineUsers int    `json:"online_users"`
	DanmuCount  int    `json:"danmu_count"`
}

// SearchAllRooms 获取房间列表
func SearchAllRooms() []room {
	var rooms []room
	global.DB.Table("danmu_rooms").Select([]string{"id", "room_name", "room_desc", "online_users, danmu_count"}).Scan(&rooms)
	return rooms
}

// SearchUser 跟用户相关的查询操作
func SearchUser(mode string, arg interface{}) (data model.User, err error) {
	switch mode {
	case "ByPass":
		{
			L := arg.(model.User)
			err = global.DB.Where("username = ? AND password = ?", L.Username, utils.MD5V([]byte(L.Password))).First(&data).Error
		}
	case "ByName":
		{
			L := arg.(string)
			err = global.DB.Where("username = ?", L).First(&data).Error
		}
	case "All":
		{
			err = global.DB.Table("users").Select("*").Scan(&data).Error
		}
	case "Existed":
		{
			L := arg.(string)
			err = global.DB.Where("username = ?", L).First(&data).Error
		}
	}
	return data, err
}
