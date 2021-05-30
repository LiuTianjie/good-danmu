/*
 * @Descripttion: your project
 * @version: 1.0
 * @Author: Nickname4th
 * @Date: 2021-05-22 16:55:05
 * @LastEditors: Nickname4th
 * @LastEditTime: 2021-05-28 17:31:17
 */
package init

import (
	"fmt"
	"good-danmu/src/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	config := "config.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置文件读取错误"))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.CONFIG.JWT); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
