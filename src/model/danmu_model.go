/*
@Time : 2021/5/31 11:40
@Author : nickname4th
@File : danmu_model
@Software: GoLand
*/
package model

import "gorm.io/gorm"

// 弹幕房间与弹幕内容是一对多关系，一个房间可以对应多条弹幕，一条弹幕只对应一个房间
// 用户和弹幕内容是多对一关系，一个用户可对应多条弹幕，一条弹幕只对应一个用户
// 用户和弹幕房间是多对多关系，一个用户可属于多个房间，一个房间可拥有多个用户

type DanmuRoom struct {
	gorm.Model
	DanmuContents []DanmuContent `json:"danmu_contents" gorm:"foreignKey:RoomId;"`
	RoomName      string         `json:"room_name"`
	RoomDesc      string         `json:"room_desc"`
	Users         []User         `json:"users" gorm:"many2many:danmu_room_users;"`
	OnlineUsers   int64          `json:"online_users" gorm:"comment:当前房间在线人数"`
	DanmuCount    int64          `json:"danmu_count"`
}

type DanmuContent struct {
	gorm.Model
	Username string `json:"username" gorm:"comment:用户名"`
	RoomId   string `json:"room_id" gorm:"comment:弹幕房间id"`
	Content  []byte `json:"content" gorm:"comment:弹幕内容"`
	Type     string `json:"type" gorm:"comment:弹幕类型"`
}
