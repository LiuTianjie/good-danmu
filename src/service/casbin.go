/*Package service
@Time : 2021/6/14 16:09
@Author : nickname4th
@File : casbin
@Software: GoLand
*/
package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"good-danmu/src/global"
	"log"
	"strings"
	"sync"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
	err            error
)

func Casbin() (*casbin.SyncedEnforcer, error) {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DB)
		if syncedEnforcer, err = casbin.NewSyncedEnforcer(global.CONFIG.Casbin.ModelPath, a); err != nil {
			log.Println("读取配置文件失败", err)
		} else {
			syncedEnforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
		}
	})
	err = syncedEnforcer.LoadPolicy()
	log.Println("LoadError", syncedEnforcer.GetPolicy())

	return syncedEnforcer, err
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return ParamsMatch(name1, name2), nil
}

// UpdateCasbin TODO: functions to add, drop rules
func UpdateCasbin() {

}
