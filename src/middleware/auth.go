/*
 * @Descripttion: auth middleware
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 16:54:00
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-28 17:53:06
 */
package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"good-danmu/src/global"
	"good-danmu/src/model"
	"good-danmu/src/utils"
	"log"
	"strconv"
	"time"
)

type JWT struct {
	SignKey []byte
}

var (
	TokenExpired     = errors.New("Token已过期")
	TokenNotValidYet = errors.New("Token未生效")
	TokenMalformed   = errors.New("Token 非法")
	TokenInvalid     = errors.New("无法解析Token")
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			token = c.Request.Header.Get("Sec-WebSocket-Protocol")
		}
		log.Println("token is ", token)
		if token == "" {
			utils.FailedMsg(401, "未授权的访问", c)
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseJWT(token)
		if err != nil {
			if err == TokenExpired {
				utils.FailedMsg(401, "授权已过期", c)
				c.Abort()
				return
			}
			utils.FailedMsg(401, "认证失败", c)
			c.Abort()
			return
		}
		var u model.User
		if err = global.DB.Preload("Role").Where("`uuid` = ?", claims.UUID.String()).First(&u).Error; err != nil {
			utils.FailedMsg(401, "认证失败", c)
			c.Abort()
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateJWT(*claims)
			newClaims, _ := j.ParseJWT(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		log.Println("user role info is:", u.UserRoleId, u.Role.RoleName)
		c.Set("claims", claims)
		c.Next()
	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.JWT.SignKey),
	}
}

func (j *JWT) CreateJWT(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

func (j *JWT) ParseJWT(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
