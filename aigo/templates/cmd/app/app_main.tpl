package main

import (
	"log"

	"gen-templates/configs"
	"gen-templates/init/api"
	"gen-templates/init/database"
)

func main() {
	conf := configs.LoadConf()
	gdb := database.InitGROMClientForMySQL(conf)
	// 1.创建路由
	r, err := api.InitGinRouters(gdb)
	if err != nil {
		log.Panic("router init error, ", err)
	}
	// Run("里面不指定端口号默认为8080")
	err = r.Run(conf.Server.Port)
	if err != nil {
		panic(err)
	}
}

