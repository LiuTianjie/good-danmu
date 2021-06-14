/*Package model
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
	UUID       uuid.UUID `json:"uuid" gorm:"index:,unique"`
	Username   string    `json:"username" gorm:"comment:用户名"`
	Password   string    `json:"password" gorm:"comment:用户密码"`
	Role       Role      `json:"role" gorm:"foreignKey:RoleId;references:UserRoleId;comment:用户角色"`
	UserRoleId string    `json:"user_ole_id" gorm:"index;comment:用户角色ID"`
}
