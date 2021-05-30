/*
 * @Descripttion: user model
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-28 15:48:45
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-28 15:50:56
 */
package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username string    `json:"username" grom:"comment:用户名"`
	Password string    `json:"password" grom:"comment:用户密码`
}
