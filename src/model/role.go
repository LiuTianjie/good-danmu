/*Package model
@Time : 2021/6/12 12:28
@Author : nickname4th
@File : role.go
@Software: GoLand
*/
package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleId   string `json:"role_id" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	RoleName string `json:"role_name"`
}

type CasbinModel struct {
	Ptype    string `json:"p_type" gorm:"column:ptype"`
	RoleName string `json:"role_name" gorm:"column:v0"`
	Path     string `json:"path" gorm:"column:v1"`
	Method   string `json:"method" gorm:"column:v2"`
}
