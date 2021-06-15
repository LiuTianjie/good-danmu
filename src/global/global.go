/*Package global
 * @Descripttion: Global structure, use viper to init.
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-28 15:53:43
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-28 16:56:13
 */
package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	DB     *gorm.DB
	RDB    *RedisDB
	CONFIG Server
	VP     *viper.Viper
)

type Server struct {
	JWT        JWT           `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	ServerInfo string        `mapstructure:"server-info" json:"server-info" yaml:"server-info"`
	Mysql      Mysql         `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis      RedisDbConfig `mapstructure:"redis" json:"redis" yaml:"redis"`
	Casbin     Casbin        `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}

type JWT struct {
	SignKey     string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"` // 服务器地址:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

type RedisDB struct {
	Rdb     *redis.ClusterClient
	AofChan chan bool
}

type RedisDbConfig struct {
	Addr     []string
	Password string
	DB       int
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

// Persistence
// Save the values per 6 minutes.
func (rdb *RedisDB) Persistence() {
	var err error
	for {
		if err = rdb.Rdb.Get("Save").Err(); err != nil {
			rdb.Rdb.Set("Save", "wait", 600*time.Second)
		} else {
			rdb.Rdb.BgSave()
			// Save data to Mysql
			time.Sleep(60 * time.Second)
		}
	}
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"` // 存放casbin模型的相对路径
}
