package main

import (
	"log"

	"demo_moudle/configs"
	"demo_moudle/init/api"
)

func main() {
	conf := configs.LoadConf()
	// 1.创建路由
	r, err := api.InitGinRouters()
	if err != nil {
		log.Panic("router init error, ", err)
	}
	// Run("里面不指定端口号默认为8080")
	err = r.Run(conf.Server.Port, conf.Server.Port)
	if err != nil {
		panic(err)
	}
}
