package convertor

import "github.com/wvegtre/gogen-cli/convertor/config"

type ConnectOperator interface {
	GenStructByDBFields(parameter GenDBCodeParameter) error
	GenServiceForDBStruct() error
	GenServerFunc() error
}

func NewMySQLDBConnect(config *config.GenConfig) ConnectOperator {
	return &MySQLDBConnect{
		Config: config,
	}
}
