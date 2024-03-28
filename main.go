package main

import (
	"boot/boot"
	"boot/conf"
	"boot/starter"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("conf.ini")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	conf := &conf.Config{}

	err = viper.Unmarshal(conf)
	if err != nil {
		panic(err)
	}

	boot.Register(&starter.LogStarter{})
	boot.Register(&starter.WebServerStarter{})
	//starter.Register(&starter.MysqlStater{})
	//starter.Register(&starter.RedisStater{})

	// 1、实例化引导程序
	boot := boot.NewBootApplication(conf)

	// 3、启动
	boot.Start()
}
