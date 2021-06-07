/*
 * @Descripttion: response
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-28 16:03:39
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-28 16:14:12
 */
package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"message"`
	Reason string      `json:"reason"`
}

func Result(code int, data interface{}, msg string, reason string, c *gin.Context) {
	c.JSON(code, Response{
		code,
		data,
		msg,
		reason,
	})
}

// 统一封装正确或错误请求处理

func Ok(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "SUCCESS", "操作成功", c)
}

func OkMsg(code int, reason string, c *gin.Context) {
	Result(code, map[string]interface{}{}, "SUCCESS", reason, c)
}

func OkDetail(code int, data interface{}, reason string, c *gin.Context) {
	Result(code, data, "SUCCESS", reason, c)
}

func Failed(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "FAILED", "操作失败", c)
}

func FailedMsg(code int, reason string, c *gin.Context) {
	Result(code, map[string]interface{}{}, "FAILED", reason, c)
}

func FailedDetail(code int, data interface{}, reason string, c *gin.Context) {
	Result(code, data, "FAILED", reason, c)
}
