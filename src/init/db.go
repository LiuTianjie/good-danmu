/*
 * @Descripttion: db models
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 16:54:30
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-28 16:59:55
 */
package init

import (
	"github.com/go-redis/redis"
	"good-danmu/src/global"
	"good-danmu/src/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Println("MySql启动错误")
		return nil
	}
	return db
}

func Gorm(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.DanmuRoom{},
		model.DanmuContent{},
	)
	if err != nil {
		log.Println("初始化数据库失败")
		os.Exit(0)
	}
	log.Println("初始化数据库成功")
}

func ConnectRDB() *global.RedisDB {
	r := global.CONFIG.Redis
	rDb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    r.Addr,
		Password: r.Password,
		//DB:       r.DB,
	})
	redisDb := &global.RedisDB{
		Rdb:     rDb,
		AofChan: make(chan bool, 1),
	}
	if pong, err := redisDb.Rdb.Ping().Result(); err != nil {
		log.Println("Redis连接失败！")
	} else {
		log.Println("Redis连接成功，pong is :", pong, err)
	}
	return redisDb
}
