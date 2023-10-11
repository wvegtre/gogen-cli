package main

import (
	"log"

	"chocolate/init/api"
	"chocolate/middleware/server_ctx"
)

func main() {
	server_ctx.Init()
	// 1.创建路由
	r, err := api.InitGinRouters()
	if err != nil {
		log.Panic("router init error, ", err)
	}
	// Run("里面不指定端口号默认为8080")
	err = r.Run(server_ctx.Get().ServerConfigs.Server.Port)
	if err != nil {
		panic(err)
	}
}

