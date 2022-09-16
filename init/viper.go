package init

import (
	"github.com/Armove/DouYinProject/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func Viper() {
	//
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yml")
	//
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("获取配置文件错误")
	}
	//
	err = viper.Unmarshal(&global.CONFIG)
	if err != nil {
		log.Panic("viper反序列化错误")
	}
	log.Println(global.CONFIG)
	//
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件被修改：", e.Name)
	})
}
