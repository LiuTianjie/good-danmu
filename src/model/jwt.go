/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-28 17:44:21
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-29 10:24:56
 */
package model

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID       uuid.UUID
	ID         uint
	Username   string
	Privilege  int
	BufferTime int64
	jwt.StandardClaims
}
