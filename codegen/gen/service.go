package gen

import "github.com/wvegtre/gogen-cli/gen/config"

type ConnectOperator interface {
	GenStructByDBFields(parameter GenDBCodeParameter) error
	GenServiceForDBStruct() error
}

func NewMySQLDBConnect(config *config.GenConfig) ConnectOperator {
	return &MySQLDBConnect{
		Config: config,
	}
}
