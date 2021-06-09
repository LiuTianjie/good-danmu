/*Package handler
 * @Descripttion: danmu handler, do some pre precess.
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 17:25:24
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-23 21:12:49
 */
package handler

import (
	"encoding/json"
	"good-danmu/src/global"
	dm "good-danmu/src/service"
	"good-danmu/src/utils"
	"good-danmu/src/utils/parse"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// DanmuHandler
// 1. Check whether there is a token.
// 2. After upgrade, the conn is a real websocket connection.
// 3. We run two goroutines for it's read and write loop.
// 4 .Get the stored danmus and push it to the client, to get high performance,
// we need to zip the danmus rather to send it one by one. Here need to negotiate
// with front end developers to standardize data formats.
func DanmuHandler(c *gin.Context) {
	danmuId := c.Param("id")
	token := c.GetHeader("Sec-WebSocket-Protocol")
	var (
		conn     *websocket.Conn
		err      error
		danmu    *dm.DanmuServer
		Username string
	)
	if Username, err = parse.GetUserFromToken(token); err != nil {
		log.Println("没有找到用户")
		utils.FailedMsg(400, "没有找到对应的用户", c)
		return
	}
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{token},
	}
	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	danmu = dm.InitDanmuServer(conn, danmuId, Username)
	dm.DanmuChannels[danmuId] = append(dm.DanmuChannels[danmuId], danmu)
	existedDanmus := global.RDB.Rdb.LRange(danmuId, 0, -1)
	collection, _ := json.Marshal(existedDanmus.Val())
	danmu.OutChan <- collection
}
