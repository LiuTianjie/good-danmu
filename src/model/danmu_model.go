/*
@Time : 2021/5/31 11:40
@Author : nickname4th
@File : danmu_model
@Software: GoLand
*/
package model

import "gorm.io/gorm"

type Danmu struct {
	gorm.Model
	Username string `json:"username" gorm:"comment:用户名"`
	Content  string `json:"content" gorm:"comment:弹幕内容"`
	Type     string `json:"type" gorm:"comment:弹幕类型"`
}
