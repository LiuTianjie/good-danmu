/*Package middleware
@Time : 2021/6/13 11:00
@Author : nickname4th
@File : role_judge
@Software: GoLand
*/
package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"good-danmu/src/model"
	"good-danmu/src/service"
	"good-danmu/src/utils"
	"log"
)

func RoleJudge(c *gin.Context) {
	var (
		e   *casbin.SyncedEnforcer
		err error
	)
	claims, _ := c.Get("claims")
	waitUse := claims.(*model.CustomClaims)
	sub := waitUse.RoleId
	obj := c.Request.URL.RequestURI()
	act := c.Request.Method
	if e, err = service.Casbin(); err != nil {
		log.Println("读取配置文件失败")
		utils.FailedMsg(400, "权限不足", c)
		c.Abort()
		return
	}
	success, _ := e.Enforce(sub, obj, act)
	if success {
		c.Next()
	} else {
		utils.FailedMsg(400, "权限不足", c)
		c.Abort()
	}
}
