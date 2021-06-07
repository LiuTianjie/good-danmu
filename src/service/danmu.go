/*
 * @Descripttion: Define danmu struct
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 17:48:14
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-23 21:39:13
 */
package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"good-danmu/src/global"
	"good-danmu/src/model"
	"log"
	"sync"
)

// TODO: when client is offline, remove it from the DanmuServer list.
// Since each time we establish a connection, the size of the list will expand,
// and finally, it will be a nightmare.
// Think about two ways:
// 1. Since we must transverse the server list to broadcast the message
// to all connections, we can remove the danmuServer which 'isClosed' filed is 'false'.
// While there may be a concurrent error because each client may send message at the same time.
// So consider about using sync mutex.
// 2. Use a channel to notify whether one can use the data structure.
//TODO:Use Redis to replace the DanmuChannels
var DanmuChannels = map[string][]*DanmuServer{}

type DanmuServer struct {
	dmName    string
	Username  string
	uid       uuid.UUID
	conn      *websocket.Conn
	InChan    chan []byte
	OutChan   chan []byte
	CloseChan chan byte
	isClosed  bool
	mutex     sync.Mutex
}

func (dm *DanmuServer) Read() (data []byte, err error) {
	select {
	case data = <-dm.InChan:
	case <-dm.CloseChan:
		{
			err = errors.New("danmu connection is closed")
		}
	}
	return
}

func (dm *DanmuServer) Write(data []byte) (err error) {
	select {
	case dm.OutChan <- data:
	case <-dm.CloseChan:
		{
			err = errors.New("danmu connection is closed")
		}
	}
	return
}

func (dm *DanmuServer) Close() {
	log.Println(dm.Username, "'s connection close")
	var err error
	if err = dm.conn.Close(); err != nil {
		log.Println("关闭出错")
	}
	dm.mutex.Lock()
	if !dm.isClosed {
		close(dm.CloseChan)
		dm.isClosed = true
	}
	dm.mutex.Unlock()
}

type Msg struct {
	Time    string `json:"time"`
	Content string `json:"content"`
}

func (dm *DanmuServer) ReadLoop() {
	log.Println(dm.dmName + "is reading looply~")
	var (
		err  error
		data []byte
	)
	for {
		if _, data, err = dm.conn.ReadMessage(); err != nil {
			goto ERR
		}
		var msg Msg
		err = json.Unmarshal(data, &msg)
		select {
		case dm.InChan <- []byte(msg.Content):
			{
				//SaveDanmu(dm.dmName, dm.Username, data)
				global.RDB.Rdb.Set(string(dm.uid[:])+msg.Time, data, -1)
				// Traverse the channel in the server
				// TODO: Use Redis connection pool to broadcast the message.
				for _, v := range DanmuChannels[dm.dmName] {
					if v.uid != dm.uid && !v.isClosed {
						v.InChan <- []byte(msg.Content)
						fmt.Println("send to:", v.uid)
					}
				}
			}
		case <-dm.CloseChan:
			{
				goto ERR
			}
		}
	}
ERR:
	dm.Close()
}

func (dm *DanmuServer) WriteLoop() {
	log.Println(dm.dmName + "is writing looply~")
	var (
		err  error
		data []byte
	)
	for {
		select {
		case data = <-dm.OutChan:
		case <-dm.CloseChan:
			{
				goto ERR
			}
		}
		if err = dm.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	dm.Close()
}

func InitDanmuServer(wsConn *websocket.Conn, dmName string, Username string) (conn *DanmuServer, err error) {
	log.Println(Username + " Joined " + dmName)
	conn = &DanmuServer{
		dmName:   dmName,
		Username: Username,
		uid:      uuid.NewV4(),
		conn:     wsConn,
		// Use 1000 byte to cache the message
		InChan:    make(chan []byte, 1000),
		OutChan:   make(chan []byte, 1000),
		CloseChan: make(chan byte, 1),
	}
	go conn.ReadLoop()
	go conn.WriteLoop()
	return
}

func SaveDanmu(DanmuId, Username string, DanmuContent []byte) {
	danmu := &model.DanmuContent{
		Content:  DanmuContent,
		RoomId:   DanmuId,
		Username: Username,
		Type:     "normal",
	}
	err := global.DB.Create(&danmu).Error
	if err != nil {
		log.Println(err)
		log.Println("插入出错")
	} else {
		log.Println("插入成功！")
	}
}
