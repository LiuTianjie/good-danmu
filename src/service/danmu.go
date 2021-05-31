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
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
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
var DanmuChannels = map[string][]*DanmuServer{}

type DanmuServer struct {
	dmName    string
	Username  string
	uid       uuid.UUID
	conn      *websocket.Conn
	InChan    chan []byte
	OutChan   chan []byte
	CloseChan chan byte
	isCLosed  bool
	mutex     sync.Mutex
}

func (dm *DanmuServer) Read() (data []byte, err error) {
	select {
	case data = <-dm.InChan:
	case <-dm.CloseChan:
		{
			err = errors.New("danmu connection is closed!")
		}
	}
	return
}

func (dm *DanmuServer) Write(data []byte) (err error) {
	select {
	case dm.OutChan <- data:
	case <-dm.CloseChan:
		{
			err = errors.New("danmu connection is closed!")
		}
	}
	return
}

func (dm *DanmuServer) Close() {
	dm.conn.Close()
	dm.mutex.Lock()
	if !dm.isCLosed {
		close(dm.CloseChan)
		dm.isCLosed = true
	}
	dm.mutex.Unlock()
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
		select {
		case dm.InChan <- data:
			{
				log.Println(DanmuChannels)
				// traverse the channel in the server
				for _, v := range DanmuChannels[dm.dmName] {
					if v.uid != dm.uid && !v.isCLosed {
						v.InChan <- data
						fmt.Println("send to:", v.uid)
					}
					log.Println(dm.Username, "said:", data)
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
		if dm.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	dm.Close()
}

func InitDanmuServer(wsConn *websocket.Conn, dmName string, Username string) (conn *DanmuServer, err error) {
	conn = &DanmuServer{
		dmName:    dmName,
		Username:  Username,
		uid:       uuid.NewV4(),
		conn:      wsConn,
		InChan:    make(chan []byte, 1000),
		OutChan:   make(chan []byte, 1000),
		CloseChan: make(chan byte, 1),
	}
	go conn.ReadLoop()
	go conn.WriteLoop()
	return
}
