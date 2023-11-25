package main

import (
	"chocolate/init/api"
	_ "chocolate/middleware/log"
	"chocolate/middleware/server_ctx"
)

func main() {
	server_ctx.Init()
	// 1.创建路由
	r, err := api.InitGinRouters()
	if err != nil {
	   panic(err)
	}
	// Run("里面不指定端口号默认为8080")
	err = r.Run(server_ctx.Get().ServerConfigs.Server.Port)
	if err != nil {
		panic(err)
	}
}

