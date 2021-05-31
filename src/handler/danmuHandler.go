/*
 * @Descripttion: danmu handler
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 17:25:24
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-23 21:12:49
 */
package handler

import (
	dm "good-danmu/src/service"
	"good-danmu/src/utils"
	"good-danmu/src/utils/parse"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func DanmuHandler(c *gin.Context) {
	danmuId := c.Param("id")
	token := c.GetHeader("Sec-WebSocket-Protocol")
	var (
		conn     *websocket.Conn
		err      error
		danmu    *dm.DanmuServer
		data     []byte
		Username string
	)
	if Username, err = parse.GetUserFromToken(token); err != nil {
		utils.FailedMsg(400, "没有找到对应的用户", c)
		return
	}
	log.Println(Username + "joined" + danmuId)
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 校验是否携带token
		Subprotocols: []string{token},
	}
	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	// After upgrade, the conn is a real websocket connection.
	// After InitDanmuServer function is executed, we run two goroutines for it's read and write loop.
	if danmu, err = dm.InitDanmuServer(conn, danmuId, Username); err != nil {
		goto ERR
	}
	dm.DanmuChannels[danmuId] = append(dm.DanmuChannels[danmuId], danmu)
	for {
		if data, err = danmu.Read(); err != nil {
			goto ERR
		}
		if err = danmu.Write(data); err != nil {
			goto ERR
		}
	}

ERR:
	danmu.Close()
}
