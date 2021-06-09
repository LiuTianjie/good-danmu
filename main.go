/*
 * @Descripttion: good-danmu
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 16:47:46
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-28 16:55:45
 */
package main

import (
	"good-danmu/src/global"
	in "good-danmu/src/init"
)

func main() {
	global.VP = in.Viper()
	global.DB = in.ConnectDB()
	if global.DB != nil {
		in.Gorm(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}
	global.RDB = in.ConnectRDB()
	// Use AOF instead of Rdb, the process need to config in the server end.
	//go global.RDB.Persistence()
	router := in.Routers()
	router.Run(":8000")
}
