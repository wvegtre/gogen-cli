package server_ctx

import (
	"sync"

	"chocolate/configs"
)

var _serverCtx ServerContext

type ServerContext struct {
	ServerConfigs    configs.Conf
	ServerComponents *Components
}

func Get() ServerContext {
	return _serverCtx
}

func Init() {
	once := sync.Once{}
	once.Do(func() {
		serverConfigs := configs.LoadConf()
		serverComponents := initComponents(serverConfigs)
		_serverCtx = ServerContext{
			ServerConfigs:    serverConfigs,
			ServerComponents: serverComponents,
		}
	})
}
