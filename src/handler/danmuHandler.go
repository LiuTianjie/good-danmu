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
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func DanmuHandler(c *gin.Context) {

	danmuId := c.Param("id")
	log.Println("enter handler:", danmuId)
	log.Println(danmuId)
	var (
		conn  *websocket.Conn
		err   error
		danmu *dm.DanmuServer
		data  []byte
	)

	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	// After upgrade, the conn is a real websocket connection.
	if danmu, err = dm.InitDanmuServer(conn, danmuId); err != nil {
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
