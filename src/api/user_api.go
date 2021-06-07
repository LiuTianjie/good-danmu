/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-28 15:47:36
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-29 10:53:56
 */
package api

import (
	"encoding/json"
	"errors"
	"good-danmu/src/global"
	h "good-danmu/src/handler"
	"good-danmu/src/middleware"
	"good-danmu/src/model"
	"good-danmu/src/utils"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type UserInfo struct {
	Token     string    `json:"token"`
	UserName  string    `json:"user_name"`
	UserId    uuid.UUID `json:"user_id"`
	Privilege int       `json:"privilege"`
}

func Register(c *gin.Context) {
	var (
		u   model.User
		err error
	)
	if err = c.ShouldBindJSON(&u); err != nil {
		utils.FailedMsg(400, "注册信息有误", c)
		return
	}
	if _, err = h.SearchUser("Existed", u.Username); !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.FailedMsg(400, "用户名已注册", c)
	} else {
		user := &model.User{Username: u.Username, Password: u.Password}
		user.Password = utils.MD5V([]byte(user.Password))
		user.UUID = uuid.NewV4()
		if err = global.DB.Create(&user).Error; err != nil {
			utils.FailedMsg(400, "注册失败", c)
		} else {
			utils.OkMsg(200, "注册成功", c)
		}
	}
}

func Login(c *gin.Context) {
	var (
		err  error
		L    model.User
		user model.User
	)
	if err = c.ShouldBind(&L); err != nil {
		utils.FailedMsg(400, "用户名/密码不能为空", c)
		c.Abort()
	} else {
		if user, err = h.SearchUser("ByPass", L); err != nil {
			utils.FailedMsg(401, "用户名/密码/错误", c)
			c.Abort()
		} else {
			tokenNext(c, user)
		}
	}
}

func tokenNext(c *gin.Context, user model.User) {
	var (
		err       error
		token     string
		tokenInfo UserInfo
	)
	// 先查redis中有没有token
	if token, err = global.RDB.Get(user.Username).Result(); err != nil {
		log.Println("Redis中没有token")
	} else {
		if err = json.Unmarshal([]byte(token), &tokenInfo); err != nil {
			utils.FailedMsg(400, "获取token失败", c)
			return
		}
		// 在Redis中查找到，返回tokenInfo
		utils.OkDetail(200, tokenInfo, "登录成功", c)
		return
	}
	// 在Redis中没有查找到，则签发token
	j := middleware.JWT{
		SignKey: []byte(global.CONFIG.JWT.SignKey),
	}
	claims := model.CustomClaims{
		UUID:       user.UUID,
		Username:   user.Username,
		BufferTime: global.CONFIG.JWT.BufferTime,
		Privilege:  1,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime,
			Issuer:    global.CONFIG.JWT.SignKey,
		},
	}
	if token, err = j.CreateJWT(claims); err != nil {
		utils.FailedMsg(400, "获取token失败", c)
		return
	} else {
		// 将新签发的token写入redis
		tokenInfo = UserInfo{
			token,
			user.Username,
			user.UUID,
			claims.Privilege,
		}
		data, _ := json.Marshal(tokenInfo)
		utils.OkDetail(200, tokenInfo, "登录成功", c)
		if err = global.RDB.Set(user.Username, data, 100*time.Second).Err(); err != nil {
			log.Println("向Redis存储tokenInfo的过程出错")
			return
		}
	}
}

func GetUserInfo(c *gin.Context) {
	var (
		err      error
		userName string
		user     model.User
	)
	if err = c.ShouldBind(&userName); err != nil {
		utils.FailedMsg(404, "参数缺失", c)
		c.Abort()
	} else {
		if user, err = h.SearchUser("ByName", userName); err != nil {
			utils.OkMsg(200, "无此用户", c)
			return
		} else {
			utils.OkDetail(200, &UserInfo{UserName: user.Username, UserId: user.UUID}, "查找成功", c)
			return
		}
	}
}
